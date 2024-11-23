<script setup lang="ts">
import { cn } from "@/lib/utils";
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";
import { toDate } from "radix-vue/date";
import {
  CalendarDate,
  DateFormatter,
  getLocalTimeZone,
  parseDate,
  today,
} from "@internationalized/date";
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
import {
  Select,
  SelectContent,
  SelectGroup,
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
import { Textarea } from "@/components/ui/textarea";
import { Input } from "@/components/ui/input";

const props = defineProps<{
  type: string;
  data: any[];
}>();

/* handle form */
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
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const formData = new FormData();
    Object.entries(values).forEach(([key, value]) => {
      formData.append(key, value);
    });

    const response = await $fetch(
      `${useRuntimeConfig().public.apiBase}/asset-aplikasi`,
      {
        method: "POST",
        body: formData,
      }
    );

    console.log("Data berhasil disimpan:", response);
    navigateTo("/aplikasi");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
});
</script>

<template>
  <div
    class="p-8 bg-white shadow-lg rounded-lg"
    :class="isOpen ? 'lg:ml-64' : ''"
  >
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
                    {{
                      values.tanggalPembuatan
                        ? values.tanggalPembuatan
                        : "Pilih tanggal"
                    }}
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
                    {{
                      values.tanggalTerima
                        ? values.tanggalTerima
                        : "Pilih tanggal"
                    }}
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
                    {{
                      values.tanggalAktifDomain
                        ? values.tanggalAktifDomain
                        : "Pilih tanggal"
                    }}
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
                    {{
                      values.tanggalExpiredDomain
                        ? values.tanggalExpiredDomain
                        : "Pilih tanggal"
                    }}
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
        <Button variant="outline" @click="$router.back()">Cancel</Button>
        <Button type="submit">Submit</Button>
      </div>
    </Form>
  </div>
</template>
