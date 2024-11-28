<script setup lang="ts">
import { cn } from "@/lib/utils";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { CalendarIcon } from "@radix-icons/vue";
import { format } from "date-fns";
import { ref } from "vue";

import { Button } from "@/components/ui/button";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Input } from "@/components/ui/input";

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    employeeId: z.string().min(1, { message: "ID Pegawai wajib diisi" }),
    name: z.string().min(1, { message: "Nama wajib diisi" }),
    email: z.string().email({ message: "Email tidak valid" }),
    position: z.string().min(1, { message: "Jabatan wajib diisi" }),
    division: z.string().min(1, { message: "Divisi wajib diisi" }),
    joinDate: z.string().min(1, { message: "Tanggal bergabung wajib diisi" }),
  })
);

const formData = ref({
  employeeId: "",
  name: "",
  email: "",
  position: "",
  division: "",
  joinDate: "",
});

const selectedDate = ref();
const errors = ref({});
const isSubmitting = ref(false);

const validateForm = () => {
  errors.value = {};

  if (!formData.value.employeeId) {
    errors.value.employeeId = "ID Pegawai wajib diisi";
  }
  if (!formData.value.name) {
    errors.value.name = "Nama wajib diisi";
  }
  if (!formData.value.email) {
    errors.value.email = "Email wajib diisi";
  } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.value.email)) {
    errors.value.email = "Format email tidak valid";
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

const onSubmit = async () => {
  if (!validateForm()) {
    return;
  }

  try {
    isSubmitting.value = true;

    const response = await $fetch("http://103.127.139.11:5000/api/user", {
      method: "POST",
      body: {
        employee_id: formData.value.employeeId,
        name: formData.value.name,
        email: formData.value.email,
        position: formData.value.position,
        division: formData.value.division,
        join_date: formData.value.joinDate,
      },
    });

    console.log("Data berhasil disimpan:", response);
    navigateTo("/employee");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
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
            <FormField name="employeeId">
              <FormItem>
                <FormLabel>ID Pegawai</FormLabel>
                <FormControl>
                  <Input
                    v-model="formData.employeeId"
                    placeholder="Masukkan ID Pegawai"
                    :class="{ 'border-red-500': errors.employeeId }"
                  />
                </FormControl>
                <FormMessage v-if="errors.employeeId" class="text-red-500">
                  {{ errors.employeeId }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="name">
              <FormItem>
                <FormLabel>Nama</FormLabel>
                <FormControl>
                  <Input
                    v-model="formData.name"
                    placeholder="Masukkan Nama Lengkap"
                    :class="{ 'border-red-500': errors.name }"
                  />
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
                  <Input
                    type="email"
                    v-model="formData.email"
                    placeholder="Masukkan Email"
                    :class="{ 'border-red-500': errors.email }"
                  />
                </FormControl>
                <FormMessage v-if="errors.email" class="text-red-500">
                  {{ errors.email }}
                </FormMessage>
              </FormItem>
            </FormField>

            <FormField name="position">
              <FormItem>
                <FormLabel>Jabatan</FormLabel>
                <FormControl>
                  <Select
                    v-model="formData.position"
                    :class="{ 'border-red-500': errors.position }"
                  >
                    <SelectTrigger>
                      <SelectValue placeholder="Pilih Jabatan" />
                    </SelectTrigger>
                    <SelectContent>
                      <SelectGroup>
                        <SelectItem value="AD">Admin</SelectItem>
                        <SelectItem value="SV">Supervisor</SelectItem>
                        <SelectItem value="ST">Staff</SelectItem>
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
                  <Input
                    v-model="formData.division"
                    placeholder="Masukkan Divisi"
                    :class="{ 'border-red-500': errors.division }"
                  />
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
                      <Button
                        variant="outline"
                        class="w-full justify-start text-left font-normal"
                        :class="{ 'border-red-500': errors.joinDate }"
                      >
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {{
                          selectedDate
                            ? format(selectedDate, "dd/MM/yyyy")
                            : "Pilih tanggal"
                        }}
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent>
                      <Calendar
                        mode="single"
                        v-model="selectedDate"
                        @update:model-value="updateDate"
                      />
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
