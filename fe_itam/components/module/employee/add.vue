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
            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="employeeId" class="text-right">ID Pegawai</Label>
              <Input
                id="employeeId"
                v-model="form.employeeId"
                placeholder="Masukkan ID Pegawai"
                class="col-span-3"
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="name" class="text-right">Nama</Label>
              <Input
                id="name"
                v-model="form.name"
                placeholder="Masukkan Nama Lengkap"
                class="col-span-3"
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="email" class="text-right">Email</Label>
              <Input
                id="email"
                type="email"
                v-model="form.email"
                placeholder="Masukkan Email"
                class="col-span-3"
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="position" class="text-right">Jabatan</Label>
              <Input
                id="position"
                v-model="form.position"
                placeholder="Masukkan Jabatan"
                class="col-span-3"
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="division" class="text-right">Divisi</Label>
              <Input
                id="division"
                v-model="form.division"
                placeholder="Masukkan Divisi"
                class="col-span-3"
              />
            </div>

            <div class="grid grid-cols-4 items-center gap-4">
              <Label for="joinDate" class="text-right">Tanggal Bergabung</Label>
              <Input
                id="joinDate"
                type="date"
                v-model="form.joinDate"
                class="col-span-3"
              />
            </div>
          </div>

          <div class="flex justify-end">
            <Button type="submit">Simpan</Button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

const form = ref({
  employeeId: "",
  name: "",
  email: "",
  position: "",
  division: "",
  joinDate: "",
});

const onSubmit = () => {
  // Implementasi logika penyimpanan data
  console.log("Form submitted:", form.value);
};
</script>
