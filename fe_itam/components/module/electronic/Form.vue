<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { useForm } from "vee-validate";
import * as z from "zod";

/**
 * Props untuk menerima tipe dan data dari parent component
 * @param {string} type - Tipe aset elektronik
 * @param {any[]} data - Data array yang diterima dari parent
 */
const props = defineProps<{
  type: string;
  data: any[];
}>();

// Form validation schema
const formSchema = z.object({
  namasupllier: z.string().min(1, "Supplier harus dipilih"),
  tipeasset: z.string().min(1, "Tipe aset harus dipilih"),
  notaNumber: z.string().min(1, "Nomor nota harus diisi"),
  kondisiasset: z.string().min(1, "Kondisi aset harus dipilih"),
  status: z.string().min(1, "Status harus dipilih"),
  lokasipenerima: z.string().min(1, "Lokasi penerima harus diisi"),
  nilaiasset: z.number().min(0, "Nilai aset harus diisi"),
  achievedDate: z.string().min(1, "Tanggal penerimaan harus diisi"),
  tandaterima: z.any(),
  hasilpemeriksaan: z.any(),
  tanggalaktivasi: z.string().min(1, "Tanggal aktivasi harus diisi"),
  penanggungjawabasset: z.string().min(1, "Penanggung jawab harus dipilih"),
  divpengguna: z.string().min(1, "Divisi pengguna harus dipilih"),
  notapembelian: z.any(),
  lokasipenyimpanan: z.string().min(1, "Lokasi penyimpanan harus diisi"),
  jangkamasapakaiaset: z.string().min(1, "Jangka masa pakai harus dipilih"),
  model: z.string().min(1, "Model harus diisi"),
  tanggalasetkeluar: z.string().min(1, "Tanggal aset keluar harus diisi"),
  nomorgaransi: z.string().min(1, "Nomor garansi harus diisi"),
  merk: z.string().min(1, "Merek harus diisi"),
  garansimulai: z.string().min(1, "Tanggal mulai garansi harus diisi"),
  garansiakhir: z.string().min(1, "Tanggal akhir garansi harus diisi"),
  serialnumber: z.string().min(1, "Serial number harus diisi"),
  prosesor: z.string().optional(),
  ram: z.number().optional(),
  tiperam: z.string().optional(),
  rom: z.number().optional(),
});

// Initialize form
const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    namasupllier: "",
    tipeasset: "",
    notaNumber: "",
    kondisiasset: "",
    status: "",
    lokasipenerima: "",
    nilaiasset: 0,
    achievedDate: "",
    tandaterima: null,
    hasilpemeriksaan: null,
    tanggalaktivasi: "",
    penanggungjawabasset: "",
    divpengguna: "",
    notapembelian: null,
    lokasipenyimpanan: "",
    jangkamasapakaiaset: "",
    model: "",
    tanggalasetkeluar: "",
    nomorgaransi: "",
    merk: "",
    garansimulai: "",
    garansiakhir: "",
    serialnumber: "",
    prosesor: "",
    ram: 0,
    tiperam: "",
    rom: 0,
  },
});

/**
 * Method untuk handle form submission
 * Mengirim data form ke API endpoint
 */
const onSubmit = async (values) => {
  try {
    // TODO: Implementasi POST request ke API endpoint
    // POST /api/assets dengan form data
    console.log(values);
  } catch (error) {
    console.error("Error submitting form:", error);
  }
};
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Elektronik</h1>

    <Form @submit="onSubmit" :form="form">
      <div class="grid grid-cols-2 gap-x-4">
        <!-- Form fields -->
        <FormField
          v-for="field in Object.keys(form.values)"
          :key="field"
          :name="field"
          v-slot="{ field: formField, error }"
        >
          <FormItem class="mb-4">
            <FormLabel>{{ field }}</FormLabel>
            <FormControl>
              <!-- Select fields -->
              <Select
                v-if="
                  [
                    'Nama Supllier',
                    'Tipe Asset',
                    'Kondisi Asset',
                    'Status',
                    'Penanggung Jawab Asset',
                    'Divisi Pengguna',
                    'Jangka Masa Pakai Aset',
                  ].includes(field)
                "
                v-bind="formField"
              >
                <SelectTrigger>
                  <SelectValue :placeholder="`Pilih ${field}`" />
                </SelectTrigger>
                <SelectContent>
                  <!-- Add SelectItems based on options -->
                </SelectContent>
              </Select>

              <!-- File inputs -->
              <Input
                v-else-if="
                  [
                    'tanda terima',
                    'hasil pemeriksaan',
                    'nota pembelian',
                  ].includes(field)
                "
                type="file"
                v-bind="formField"
              />

              <!-- Date inputs -->
              <Input
                v-else-if="
                  [
                    'tanggal penerimaan',
                    'tanggal aktivasi',
                    'tanggal aset keluar',
                    'tanggal mulai garansi',
                    'tanggal akhir garansi',
                  ].includes(field)
                "
                type="date"
                v-bind="formField"
              />

              <!-- Number inputs -->
              <Input
                v-else-if="['nilai aset', 'ram', 'rom'].includes(field)"
                type="number"
                v-bind="formField"
              />

              <!-- Default text input -->
              <Input v-else type="text" v-bind="formField" />
            </FormControl>
            <FormMessage>{{ error }}</FormMessage>
          </FormItem>
        </FormField>
      </div>

      <div class="flex justify-end mt-4 space-x-2">
        <Button variant="outline" type="button"> Cancel </Button>
        <Button variant="outline" type="button"> Next </Button>
        <Button type="submit"> Submit </Button>
      </div>
    </Form>
  </div>
</template>

<style scoped>
/* Tambahkan style kustom jika diperlukan */
</style>
