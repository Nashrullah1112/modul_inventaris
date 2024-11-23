<script setup lang="ts">
import { cn } from "@/lib/utils";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { CalendarIcon } from "@radix-icons/vue";

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

const { handleSubmit, values } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const response = await $fetch(
      `${useRuntimeConfig().public.apiBase}/employees`,
      {
        method: "POST",
        body: values,
      }
    );
    console.log("Data berhasil disimpan:", response);
    navigateTo("/employee");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
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
        <Form @submit="onSubmit" class="space-y-4">
          <div class="grid gap-4 py-4">
            <FormField name="employeeId">
              <FormItem>
                <FormLabel>ID Pegawai</FormLabel>
                <FormControl>
                  <Input
                    v-model="values.employeeId"
                    placeholder="Masukkan ID Pegawai"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField name="name">
              <FormItem>
                <FormLabel>Nama</FormLabel>
                <FormControl>
                  <Input
                    v-model="values.name"
                    placeholder="Masukkan Nama Lengkap"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField name="email">
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    type="email"
                    v-model="values.email"
                    placeholder="Masukkan Email"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField name="position">
              <FormItem>
                <FormLabel>Jabatan</FormLabel>
                <FormControl>
                  <Input
                    v-model="values.position"
                    placeholder="Masukkan Jabatan"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>

            <FormField name="division">
              <FormItem>
                <FormLabel>Divisi</FormLabel>
                <FormControl>
                  <Input
                    v-model="values.division"
                    placeholder="Masukkan Divisi"
                  />
                </FormControl>
                <FormMessage />
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
                      >
                        <CalendarIcon class="mr-2 h-4 w-4" />
                        {{
                          values.joinDate ? values.joinDate : "Pilih tanggal"
                        }}
                      </Button>
                    </PopoverTrigger>
                    <PopoverContent>
                      <Calendar v-model="values.joinDate" />
                    </PopoverContent>
                  </Popover>
                </FormControl>
                <FormMessage />
              </FormItem>
            </FormField>
          </div>

          <div class="flex justify-end space-x-2">
            <Button variant="outline" @click="$router.back()">Batal</Button>
            <Button type="submit">Simpan</Button>
          </div>
        </Form>
      </div>
    </div>
  </div>
</template>
