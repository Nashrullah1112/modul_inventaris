<script setup lang="ts">
import { cn } from "@/lib/utils";
import { useToast } from '@/components/ui/toast/use-toast'
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
import { Input } from "@/components/ui/input";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig()
const route = useRoute()
const { showLoading, hideLoading } = useLoading()
const { toast } = useToast()

/* hold datefield value */
const tanggalPembuatan = computed({
  get: () => values.tanggal_pembuatan ? parseDate(values.tanggal_pembuatan) : undefined,
  set: val => val,
})
const tanggalTerima = computed({
  get: () => values.tanggal_terima ? parseDate(values.tanggal_terima) : undefined,
  set: val => val,
})
const tanggalAktif = computed({
  get: () => values.tanggal_aktif ? parseDate(values.tanggal_aktif) : undefined,
  set: val => val,
})
const tanggalKadaluarsa = computed({
  get: () => values.tanggal_kadaluarsa ? parseDate(values.tanggal_kadaluarsa) : undefined,
  set: val => val,
})


/* handle form */
const formSchema = toTypedSchema(
  z.object({
    vendor_id: z.number().min(1),
    nama_aplikasi: z.string().min(1),

    tanggal_pembuatan: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tanggal_terima: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tanggal_aktif: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tanggal_kadaluarsa: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    dokumentasi: z.any(),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});

const dataId = route.params.id
const endpoint = `/asset-lisensi${props.type == 'new' ? '' : '/' + dataId}`

const onSubmit = handleSubmit(async (values) => {
  showLoading()

  try {
    const { data, status } = await useFetch(config.public.API_URL + endpoint, {
      method: props.type == 'new' ? 'POST' : 'PATCH',
      body: values,
    })

    if (status.value == 'success') {
      toast({
        title: 'Success',
        description: `Data submitted successfully`,
      })
      navigateTo('/application')
    } else {
      toast({
        title: 'Failed',
        description: `Error when submitting data`,
      })
    }
  } catch (error) {
    console.error('Error occured:', error);
  }

  hideLoading()
});
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Lisensi Software</h1>
    
    <form>
      <div class="grid grid-cols-2 gap-x-4 gap-y-2">
        <FormField v-slot="{ componentField }" name="namavendor">
          <FormItem>
            <FormLabel>Nama Vendor</FormLabel>
            <FormControl>
              <Input
                type="text" 
                id="namavendor"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_pembuatan">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Pembuatan</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button
                    variant="outline" :class="cn(
                      'ps-3 text-start font-normal',
                      !tanggalPembuatan && 'text-muted-foreground',
                    )"
                  >
                    <span>{{ tanggalPembuatan || "Pick a date" }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar
                  v-model="tanggalPembuatan"
                  calendar-label="Date of birth"
                  @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('tanggal_pembuatan', v.toString())
                    }
                    else {
                      setFieldValue('tanggal_pembuatan', undefined)
                    }
                  }"
                />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="namaaplikasi">
          <FormItem>
            <FormLabel>Nama Aplikasi</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="namaaplikasi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_terima">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Penerimaan</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button
                    variant="outline" :class="cn(
                      'ps-3 text-start font-normal',
                      !tanggalTerima && 'text-muted-foreground',
                    )"
                  >
                    <span>{{ tanggalTerima || "Pick a date" }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar
                  v-model="tanggalTerima"
                  calendar-label="Date of birth"
                  @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('tanggal_terima', v.toString())
                    }
                    else {
                      setFieldValue('tanggal_terima', undefined)
                    }
                  }"
                />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipeplatform">
          <FormItem>
            <FormLabel>Tipe Platform</FormLabel>
            <FormControl>
              <Select v-bind="componentField" required>
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Platform" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="web">Web</SelectItem>
                  <SelectItem value="mobile">Mobile</SelectItem>
                  <SelectItem value="desktop">Desktop</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="lokasiserver">
          <FormItem>
            <FormLabel>Lokasi Server Penyimpanan</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="lokasiserver"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="urlaplikasi">
          <FormItem>
            <FormLabel>URL Aplikasi</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="urlaplikasi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_aktif">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Aktif Domain</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button
                    variant="outline" :class="cn(
                      'ps-3 text-start font-normal',
                      !tanggalAktif && 'text-muted-foreground',
                    )"
                  >
                    <span>{{ tanggalAktif || "Pick a date" }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar
                  v-model="tanggalAktif"
                  calendar-label="Date of birth"
                  @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('tanggal_aktif', v.toString())
                    }
                    else {
                      setFieldValue('tanggal_aktif', undefined)
                    }
                  }"
                />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipesertifikasi">
          <FormItem>
            <FormLabel>Tipe Sertifikasi Asset</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="tipesertifikasi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_kadaluarsa">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Expired Domain</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button
                    variant="outline" :class="cn(
                      'ps-3 text-start font-normal',
                      !tanggalKadaluarsa && 'text-muted-foreground',
                    )"
                  >
                    <span>{{ tanggalKadaluarsa || "Pick a date" }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar
                  v-model="tanggalKadaluarsa"
                  calendar-label="Date of birth"
                  @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('tanggal_kadaluarsa', v.toString())
                    }
                    else {
                      setFieldValue('tanggal_kadaluarsa', undefined)
                    }
                  }"
                />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>
        
        <FormField v-slot="{ componentField }" name="serialkey">
          <FormItem>
            <FormLabel>Serial Key</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="serialkey"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="serialnumberdevice">
          <FormItem>
            <FormLabel>Serial Number Device Terpasang</FormLabel>
            <FormControl>
              <Select v-bind="componentField" required>
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Serial Number Device" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="device1">Device 1</SelectItem>
                  <SelectItem value="device2">Device 2</SelectItem>
                  <SelectItem value="device3">Device 3</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="kategorilisensi">
          <FormItem>
            <FormLabel>Kategori Lisensi</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="kategorilisensi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipelisensi">
          <FormItem>
            <FormLabel>Tipe Lisensi</FormLabel>
            <FormControl>
              <Select v-bind="componentField" required>
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Tipe Lisensi" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="perpetual">Perpetual</SelectItem>
                  <SelectItem value="subscription">Subscription</SelectItem>
                  <SelectItem value="trial">Trial</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="versilisensi">
          <FormItem>
            <FormLabel>Versi Lisensi</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="versilisensi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="estimasipemakaian">
          <FormItem>
            <FormLabel>Estimasi Pemakaian Software</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="estimasipemakaian"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="maksimumpengguna">
          <FormItem>
            <FormLabel>Maksimum Pengguna Lisensi</FormLabel>
            <FormControl>
              <Input
                type="number"
                id="maksimumpengguna"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="limitlisensi">
          <FormItem>
            <FormLabel>Limit Lisensi</FormLabel>
            <FormControl>
              <Input
                type="text"
                id="limitlisensi"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="maksimumperangkat">
          <FormItem>
            <FormLabel>Maksimum Perangkat Pengguna Lisensi</FormLabel>
            <FormControl>
              <Input
                type="number"
                id="maksimumperangkat"
                v-bind="componentField"
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="dokumentasi">
          <FormItem>
            <FormLabel >Dokumentasi</FormLabel>
            <FormControl>
              <Input
                type="file"
                id="dokumentasi"
                v-bind="componentField"
                @change=""
                required
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>
    </form>

    <div class="flex justify-end mt-4 space-x-2">
      <Button variant="outline" @click="navigateTo('/software-license')">Cancel</Button>
      <Button @click="onSubmit">Submit</Button>
    </div>
  </div>
</template>
