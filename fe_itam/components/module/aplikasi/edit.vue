<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";
import { cn } from "@/lib/utils";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { CalendarIcon } from "@radix-icons/vue";

import { Button } from "@/components/ui/button";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Input } from "@/components/ui/input";

const route = useRoute();
const router = useRouter();

const isLoading = ref(false);
const errorMessage = ref<string | null>(null);
const application = ref<Application | null>(null);

// Fetch application data by ID
async function findByIDApplication(id: number) {
  try {
    isLoading.value = true;
    errorMessage.value = null;

    const response = await $fetch<{
      message: string;
      data: Application;
      error: any;
    }>(`http://localhost:5000/api/asset-aplikasi/${id}`, {
      method: "GET",
    });

    if (response.error) {
      throw new Error(response.error);
    }

    application.value = response.data;
  } catch (error) {
    console.error("Error fetching application:", error);
    errorMessage.value = "Failed to load application data. Please try again.";
  } finally {
    isLoading.value = false;
  }
}

// Auto-call findByIDApplication when component mounts if there's an ID
onMounted(() => {
  const id = route.params.id;
  if (id) {
    findByIDApplication(Number(id));
  }
});

// Form schema validation
const formSchema = toTypedSchema(
  z.object({
    vendor: z.string().min(1),
    tanggalPembuatan: z
      .string()
      .refine((v) => v, { message: "Tanggal pembuatan wajib diisi." }),
    namaAplikasi: z.string().min(1),
    tanggalTerima: z
      .string()
      .refine((v) => v, { message: "Tanggal terima wajib diisi." }),
    tipePlatform: z.string().min(1),
    serverPenyimpanan: z.string().min(1),
    urlAplikasi: z.string().min(1),
    tanggalAktifDomain: z
      .string()
      .refine((v) => v, { message: "Tanggal aktif domain wajib diisi." }),
    tanggalExpiredDomain: z
      .string()
      .refine((v) => v, { message: "Tanggal expired domain wajib diisi." }),
    sertifikatAplikasi: z.string().min(1),
    dokumentasi: z.any(),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
  initialValues: {
    vendor: application.value?.vendor_id || "",
    tanggalPembuatan: application.value?.tanggal_pembuatan || "",
    namaAplikasi: application.value?.nama_aplikasi || "",
    tanggalTerima: application.value?.tanggal_terima || "",
    tipePlatform: application.value?.tipe_aplikasi || "",
    serverPenyimpanan: application.value?.lokasi_server_penyimpanan || "",
    urlAplikasi: application.value?.link_aplikasi || "",
    tanggalAktifDomain: application.value?.tanggal_aktif || "",
    tanggalExpiredDomain: application.value?.tanggal_kadaluarsa || "",
    sertifikatAplikasi: application.value?.sertifikasi_aplikasi || "",
    dokumentasi: null, // handle separately if needed
  },
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      formData.append(key, value);
    });

    const method = application.value ? "PUT" : "POST"; // Decide based on whether we are updating or creating
    const url = application.value
      ? `${useRuntimeConfig().public.apiBase}/asset-aplikasi/${
          application.value.id
        }`
      : `${useRuntimeConfig().public.apiBase}/asset-aplikasi`;

    const response = await $fetch(url, {
      method,
      body: formData,
    });

    console.log("Data berhasil disimpan:", response);
    router.push("/aplikasi");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
});

// Application interface for type checking
interface Application {
  id: number;
  nama_aplikasi: string;
  tanggal_pembuatan: Date;
  tanggal_terima: Date;
  lokasi_server_penyimpanan: string;
  tipe_aplikasi: string;
  link_aplikasi: string;
  sertifikasi_aplikasi: string;
  tanggal_aktif: Date;
  tanggal_kadaluarsa: Date;
  asset_id: number;
  vendor_id: number;
}
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aplikasi</h1>

    <Form @submit="onSubmit">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="vendor">
          <FormItem>
            <FormLabel>Nama Vendor</FormLabel>
            <FormControl>
              <Select v-model="values.vendor">
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
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggalPembuatan">
          <FormItem>
            <FormLabel>Tanggal Pembuatan</FormLabel>
            <FormControl>
              <Popover>
                <PopoverTrigger as-child>
                  <Button
                    variant="outline"
                    class="w-full justify-start text-left font-normal"
                  >
                    <CalendarIcon class="mr-2 h-4 w-4" />
                    {{ values.tanggalPembuatan || "Pilih tanggal" }}
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar v-model="values.tanggalPembuatan" />
                </PopoverContent>
              </Popover>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="namaAplikasi">
          <FormItem>
            <FormLabel>Nama Aplikasi</FormLabel>
            <FormControl>
              <Input v-model="values.namaAplikasi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggalTerima">
          <FormItem>
            <FormLabel>Tanggal Terima</FormLabel>
            <FormControl>
              <Popover>
                <PopoverTrigger as-child>
                  <Button
                    variant="outline"
                    class="w-full justify-start text-left font-normal"
                  >
                    <CalendarIcon class="mr-2 h-4 w-4" />
                    {{ values.tanggalTerima || "Pilih tanggal" }}
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar v-model="values.tanggalTerima" />
                </PopoverContent>
              </Popover>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tipePlatform">
          <FormItem>
            <FormLabel>Tipe Platform</FormLabel>
            <FormControl>
              <Select v-model="values.tipePlatform">
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
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="serverPenyimpanan">
          <FormItem>
            <FormLabel>Server Penyimpanan</FormLabel>
            <FormControl>
              <Input type="number" v-model="values.serverPenyimpanan" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="urlAplikasi">
          <FormItem>
            <FormLabel>URL Aplikasi</FormLabel>
            <FormControl>
              <Input v-model="values.urlAplikasi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggalAktifDomain">
          <FormItem>
            <FormLabel>Tanggal Aktif Domain</FormLabel>
            <FormControl>
              <Popover>
                <PopoverTrigger as-child>
                  <Button
                    variant="outline"
                    class="w-full justify-start text-left font-normal"
                  >
                    <CalendarIcon class="mr-2 h-4 w-4" />
                    {{ values.tanggalAktifDomain || "Pilih tanggal" }}
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar v-model="values.tanggalAktifDomain" />
                </PopoverContent>
              </Popover>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggalExpiredDomain">
          <FormItem>
            <FormLabel>Tanggal Expired Domain</FormLabel>
            <FormControl>
              <Popover>
                <PopoverTrigger as-child>
                  <Button
                    variant="outline"
                    class="w-full justify-start text-left font-normal"
                  >
                    <CalendarIcon class="mr-2 h-4 w-4" />
                    {{ values.tanggalExpiredDomain || "Pilih tanggal" }}
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar v-model="values.tanggalExpiredDomain" />
                </PopoverContent>
              </Popover>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="sertifikatAplikasi">
          <FormItem>
            <FormLabel>Sertifikat Aplikasi</FormLabel>
            <FormControl>
              <Input v-model="values.sertifikatAplikasi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="dokumentasi">
          <FormItem>
            <FormLabel>Dokumentasi</FormLabel>
            <FormControl>
              <Input
                type="file"
                @change="(e) => setFieldValue('dokumentasi', e.target.files[0])"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>

      <div class="flex justify-end mt-4 space-x-2">
        <Button variant="outline" @click="router.back()">Cancel</Button>
        <Button type="submit">Submit</Button>
      </div>
    </Form>
  </div>
</template>
