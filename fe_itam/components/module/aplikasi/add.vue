<template>
  <div
    class="p-8 bg-white shadow-lg rounded-lg"
    :class="isOpen ? 'lg:ml-64' : ''"
  >
    <h1 class="text-2xl font-bold mb-6">Registrasi Aplikasi</h1>

    <Form @submit="registerAsset">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="vendor">
          <FormItem>
            <FormLabel>Nama Vendor</FormLabel>
            <FormControl>
              <Select v-model="vendor">
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Vendor" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="vendor1">Vendor 1</SelectItem>
                  <SelectItem value="vendor2">Vendor 2</SelectItem>
                  <SelectItem value="vendor3">Vendor 3</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="tanggalPembuatan">
          <FormItem>
            <FormLabel>Tanggal Pembuatan</FormLabel>
            <FormControl>
              <Input type="date" v-model="tanggalPembuatan" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="namaAplikasi">
          <FormItem>
            <FormLabel>Nama Aplikasi</FormLabel>
            <FormControl>
              <Input type="text" v-model="namaAplikasi" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="tanggalTerima">
          <FormItem>
            <FormLabel>Tanggal Terima</FormLabel>
            <FormControl>
              <Input type="date" v-model="tanggalTerima" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="tipePlatform">
          <FormItem>
            <FormLabel>Tipe Platform</FormLabel>
            <FormControl>
              <Select v-model="tipePlatform">
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Platform" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="web">Web</SelectItem>
                  <SelectItem value="mobile">Aplikasi Mobile</SelectItem>
                  <SelectItem value="desktop">Aplikasi Desktop</SelectItem>
                  <SelectItem value="server">Aplikasi Server</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="serverPenyimpanan">
          <FormItem>
            <FormLabel>Server Penyimpanan</FormLabel>
            <FormControl>
              <Input type="number" v-model="serverPenyimpanan" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="urlAplikasi">
          <FormItem>
            <FormLabel>URL Aplikasi</FormLabel>
            <FormControl>
              <Input type="text" v-model="urlAplikasi" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="tanggalAktifDomain">
          <FormItem>
            <FormLabel>Tanggal Aktif Domain</FormLabel>
            <FormControl>
              <Input type="date" v-model="tanggalAktifDomain" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="tanggalExpiredDomain">
          <FormItem>
            <FormLabel>Tanggal Expired Domain</FormLabel>
            <FormControl>
              <Input type="date" v-model="tanggalExpiredDomain" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="sertifikatAplikasi">
          <FormItem>
            <FormLabel>Sertifikat Aplikasi</FormLabel>
            <FormControl>
              <Input type="text" v-model="sertifikatAplikasi" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="dokumentasi">
          <FormItem>
            <FormLabel>Dokumentasi</FormLabel>
            <FormControl>
              <Input type="file" @change="handleFileUpload" />
            </FormControl>
          </FormItem>
        </FormField>
      </div>

      <div class="flex justify-end mt-4 space-x-2">
        <Button variant="outline" @click="$router.back()"> Cancel </Button>
        <Button type="submit"> Submit </Button>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { useState } from "#app";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const baseURL = "http://localhost:3000/api";

const isOpen = useState("is-sidebar-open");
const router = useRouter();

const vendor = ref("");
const tanggalPembuatan = ref("");
const namaAplikasi = ref("");
const tanggalTerima = ref("");
const tipePlatform = ref("");
const serverPenyimpanan = ref(null);
const urlAplikasi = ref("");
const tanggalAktifDomain = ref("");
const tanggalExpiredDomain = ref("");
const sertifikatAplikasi = ref("");
const dokumentasi = ref(null);

const handleFileUpload = (event) => {
  dokumentasi.value = event.target.files[0];
};

const registerAsset = async () => {
  try {
    const formData = new FormData();
    formData.append("vendor", vendor.value);
    formData.append("tanggalPembuatan", tanggalPembuatan.value);
    formData.append("namaAplikasi", namaAplikasi.value);
    formData.append("tanggalTerima", tanggalTerima.value);
    formData.append("tipePlatform", tipePlatform.value);
    formData.append("serverPenyimpanan", serverPenyimpanan.value);
    formData.append("urlAplikasi", urlAplikasi.value);
    formData.append("tanggalAktifDomain", tanggalAktifDomain.value);
    formData.append("tanggalExpiredDomain", tanggalExpiredDomain.value);
    formData.append("sertifikatAplikasi", sertifikatAplikasi.value);
    formData.append("dokumentasi", dokumentasi.value);

    const response = await $fetch(`${baseURL}/asset-aplikasi`, {
      method: "POST",
      body: formData,
    });

    console.log("Data berhasil disimpan:", response);
    router.push("/aplikasi");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
};
</script>

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
