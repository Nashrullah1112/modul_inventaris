<script setup lang="ts">
import { CalendarIcon } from "@radix-icons/vue";
import { toTypedSchema } from "@vee-validate/zod";
import { format } from "date-fns";
import { ref } from "vue";
import * as z from "zod";
import { useToast } from '@/components/ui/toast/use-toast'

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    nip: z.number().min(1, { message: "ID Pegawai wajib diisi" }),
    name: z.string().min(1, { message: "Nama wajib diisi" }),
    email: z.string().email({ message: "Email tidak valid" }),
    password: z.string().min(8, { message: "Kata sandi minimal 8 karakter" }),
    position: z.number().min(1, { message: "Jabatan wajib diisi" }),
    division: z.number().min(1, { message: "Divisi wajib diisi" }),
    joinDate: z.string().min(1, { message: "Tanggal bergabung wajib diisi" }),
  })
);

const formData = ref({
  nip: null,
  name: "",
  email: "",
  password: "",
  position: null,
  division: null,
  joinDate: "",
});

const selectedDate = ref();
const errors = ref({});
const isSubmitting = ref(false);
const config = useRuntimeConfig();
const { toast } = useToast()

const divisions = ref([]);
const positions = ref([]);

const validateForm = () => {
  errors.value = {};

  if (!formData.value.nip) {
    errors.value.nip = "ID Pegawai wajib diisi";
  }
  if (!formData.value.name) {
    errors.value.name = "Nama wajib diisi";
  }
  if (!formData.value.email) {
    errors.value.email = "Email wajib diisi";
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.value.email)) {
    errors.value.email = "Format email tidak valid";
  }
  if (!formData.value.password) {
    errors.value.password = "Kata sandi wajib diisi";
  } else if (formData.value.password.length < 8) {
    errors.value.password = "Kata sandi minimal 8 karakter";
  }
  if (!formData.value.position) {
    errors.value.position = "Jabatan wajib diisi";
  }
  if (!formData.value.division) {
    errors.value.division = "Divisi wajib diisi";
  }
  if (!formData.value.joinDate) {
    errors.value.joinDate = "Tanggal bergabung wajib diisi";
  }

  return Object.keys(errors.value).length === 0;
};

const getDivisionAndPosition = async () => {
  try {
    const divisionResponse = await $fetch(config.public.API_URL + '/divisi', {
      method: 'GET',
    });

    const positionResponse = await $fetch(config.public.API_URL + '/jabatan', {
      method: 'GET',
    });

    divisions.value = divisionResponse.data;
    positions.value = positionResponse.data;

  } catch (error) {
    console.error('Terjadi kesalahan:', error instanceof Error ? error.message : error);
  }
};

const onSubmit = async () => {
  if (!validateForm()) {
    return;
  }

  try {
    isSubmitting.value = true;

    const response = await $fetch(config.public.API_URL + '/user', {
      method: "POST",
      body: {
        nip: Number(formData.value.nip),
        nama: formData.value.name,
        email: formData.value.email,
        password: formData.value.password,
        jabatan_id: Number(formData.value.position),
        divisi_id: Number(formData.value.division),
        tanggal_bergabung: `${formData.value.joinDate}T00:00:00Z`,
      },
    });

    console.log("Data berhasil disimpan:", response);
    navigateTo("/employee");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
    toast({
      title: "Error",
      description: "Gagal menyimpan data karyawan",
      variant: "destructive"
    });
  } finally {
    isSubmitting.value = false;
  }
};

const updateDate = (date: any) => {
  if (date) {
    formData.value.joinDate = format(date, "yyyy-MM-dd");
    selectedDate.value = date;
  }
};

onMounted(async () => {
  getDivisionAndPosition();
});
</script>

<template>
  <div class="container mx-auto py-10">
    <div class="flex flex-col space-y-8">
      <div>
        <h2 class="text-2xl font-bold tracking-tight">Tambah Data Karyawan</h2>
        <p class="text-muted-foreground">
          Silakan isi form berikut untuk menambahkan data karyawan baru
        </p>
      </div>

      <div class="border rounded-lg p-4">
        <form @submit.prevent="onSubmit" class="space-y-4">
          <div class="grid gap-4 py-4">
            <FormField name="nip">
              <FormItem>
                <FormLabel>ID Pegawai</FormLabel>
                <FormControl>
                  <Input v-model="formData.nip" placeholder="Masukkan ID Pegawai"
                    :class="{ 'border-red-500': errors.nip }" />
                </FormControl>
                <FormMessage v-if="errors.nip" class="text-red-500">
                  {{ errors.nip }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="name">
              <FormItem>
                <FormLabel>Nama</FormLabel>
                <FormControl>
                  <Input v-model="formData.name" placeholder="Masukkan Nama Lengkap"
                    :class="{ 'border-red-500': errors.name }" />
                </FormControl>
                <FormMessage v-if="errors.name" class="text-red-500">
                  {{ errors.name }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="email">
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input type="email" v-model="formData.email" placeholder="Masukkan Email"
                    :class="{ 'border-red-500': errors.email }" />
                </FormControl>
                <FormMessage v-if="errors.email" class="text-red-500">
                  {{ errors.email }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="password">
              <FormItem>
                <FormLabel>Kata Sandi</FormLabel>
                <FormControl>
                  <Input 
                    type="password" 
                    v-model="formData.password" 
                    placeholder="Masukkan Kata Sandi"
                    :class="{ 'border-red-500': errors.password }" 
                  />
                </FormControl>
                <FormMessage v-if="errors.password" class="text-red-500">
                  {{ errors.password }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="position">
              <FormItem>
                <FormLabel>Jabatan</FormLabel>
                <FormControl>
                  <Select v-model="formData.position" :class="{ 'border-red-500': errors.position }">
                    <SelectTrigger>
                      <SelectValue placeholder="Pilih Jabatan" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem v-for="position in positions" :key="position.id" :value="position.id">
                          {{ position.nama }} </SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage v-if="errors.position" class="text-red-500">
                  {{ errors.position }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="division">
              <FormItem>
                <FormLabel>Divisi</FormLabel>
                <FormControl>
                  <Select v-model="formData.division" :class="{ 'border-red-500': errors.division }">
                    <SelectTrigger>
                      <SelectValue placeholder="Pilih Divisi" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem v-for="division in divisions" :key="division.id" :value="division.id">
                          {{ division.nama }} </SelectItem>
                      </SelectGroup>
                    </SelectContent>
                  </Select>
                </FormControl>
                <FormMessage v-if="errors.division" class="text-red-500">
                  {{ errors.division }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="joinDate">
              <FormItem>
                <FormLabel>Tanggal Bergabung</FormLabel>
                <FormControl>
                  <Popover>
                    <PopoverTrigger as-child>
                      <Button variant="outline" class="w-full justify-start text-left font-normal"
                        :class="{ 'border-red-500': errors.joinDate }">
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {{
                          selectedDate
                            ? format(selectedDate, "dd/MM/yyyy")
                            : "Pilih tanggal"
                        }}
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent>
                      <Calendar mode="single" v-model="selectedDate" @update:model-value="updateDate" />
                    </PopoverContent>
                  </Popover>
                </FormControl>
                <FormMessage v-if="errors.joinDate" class="text-red-500">
                  {{ errors.joinDate }}
                </FormMessage>
              </FormItem>
            </FormField>
          </div>

          <div class="flex justify-end space-x-2">
            <Button variant="outline" @click="$router.back()">Batal</Button>
            <Button type="submit" :disabled="isSubmitting">
              {{ isSubmitting ? "Menyimpan..." : "Simpan" }}
            </Button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
