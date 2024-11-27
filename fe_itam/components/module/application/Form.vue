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
import { Input } from "@/components/ui/input";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig()
const route = useRoute()
const { showLoading, hideLoading } = useLoading()
const { toast } = useToast()

const df = new DateFormatter('en-US', {
  dateStyle: 'long',
})

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

/* data select vendor */
const vendors = ref([]);

const getVendorData = async () => {
  try {
    const { data, status } = await useFetch(config.public.API_URL + '/vendor');

    if (status.value == 'success' && data.value?.data?.length) {
      vendors.value = data.value.data.map(item => {
        return {
          label: item.nama_pic,
          value: item.id,
        }
      })
    }
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
}

onMounted(() => {
  const retryInterval = setInterval(() => {
    if (!vendors.value?.length)
      getVendorData()
    else
      clearInterval(retryInterval);
  }, 500);
})

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    vendor_id: z.number().min(1),
    tanggal_pembuatan: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    nama_aplikasi: z.string().min(1),
    tanggal_terima: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tipe_aplikasi: z.string().min(1),
    lokasi_server_penyimpanan: z.string().min(1),
    link_aplikasi: z.string().min(1),
    tanggal_aktif: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tanggal_kadaluarsa: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), { message: "Invalid date format, expected YYYY-MM-DD" })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    sertifikasi_aplikasi: z.string().min(1),
    dokumentasi: z.any(),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});

const dataId = route.params.id
const endpoint = `/asset-aplikasi${props.type == 'new' ? '' : '/' + dataId}`

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

let exData = <any>{}

const getExistingData = async () => {
  showLoading()

  try {
    const { data, status } = await useFetch(config.public.API_URL + endpoint);

    if (status.value == 'success' && data.value?.data) {
      exData = data.value.data
      
      setFieldValue('link_aplikasi', exData.link_aplikasi)
      setFieldValue('lokasi_server_penyimpanan', exData.lokasi_server_penyimpanan)
      setFieldValue('nama_aplikasi', exData.nama_aplikasi)
      setFieldValue('sertifikasi_aplikasi', exData.sertifikasi_aplikasi)
      setFieldValue('tanggal_aktif', transformDate(exData.tanggal_aktif))
      setFieldValue('tanggal_kadaluarsa', transformDate(exData.tanggal_kadaluarsa))
      setFieldValue('tanggal_pembuatan', transformDate(exData.tanggal_pembuatan))
      setFieldValue('tanggal_terima', transformDate(exData.tanggal_terima))
      setFieldValue('tipe_aplikasi', exData.tipe_aplikasi)
      setFieldValue('vendor_id', exData.vendor_id)
    }
  } catch (error) {
    console.error("Error occured:", error);
  }

  hideLoading()
}

if (props.type == 'edit') {
  getExistingData()
}

const transformDate = (serverDate: string) => {
  // Parse the ISO string into a CalendarDate
  const date = parseDate(serverDate?.split('T')[0]); // Extract the "YYYY-MM-DD" part

  // Format the CalendarDate as YYYY-MM-DD
  return date.toString();
}
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">
      {{ props.type == 'new' ? 'Registrasi' : 'Edit' }} Aplikasi
    </h1>

    <form>
      <div class="grid grid-cols-2 gap-x-4 gap-y-2">
        <FormField v-slot="{ componentField }" name="vendor_id">
          <FormItem>
            <FormLabel>Nama Vendor</FormLabel>
            <FormControl>
              <Select v-bind="componentField">
                <SelectTrigger>
                  <SelectValue placeholder="Pilih Vendor" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    <template v-for="item in vendors" :key="item.value">
                      <SelectItem :value="item.value">
                        {{ item.label }}
                      </SelectItem>
                    </template>
                  </SelectGroup>
                </SelectContent>
              </Select>
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

        <FormField v-slot="{ componentField }" name="nama_aplikasi">
          <FormItem>
            <FormLabel>Nama Aplikasi</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_terima">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Terima</FormLabel>
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

        <FormField v-slot="{ componentField }" name="tipe_aplikasi">
          <FormItem>
            <FormLabel>Tipe Platform</FormLabel>
            <FormControl>
              <Select v-bind="componentField">
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

        <FormField v-slot="{ componentField }" name="lokasi_server_penyimpanan">
          <FormItem>
            <FormLabel>Server Penyimpanan</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="link_aplikasi">
          <FormItem>
            <FormLabel>URL Aplikasi</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
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

        <FormField v-slot="{ componentField }" name="sertifikasi_aplikasi">
          <FormItem>
            <FormLabel>Sertifikat Aplikasi</FormLabel>
            <FormControl>
              <Input v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="dokumentasi">
          <FormItem>
            <FormLabel>Dokumentasi</FormLabel>
            <FormControl>
              <Input
                type="file"
                @change="(e) => setFieldValue('dokumentasi', e.target.files[0])"
                v-bind="componentField"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>
    </form>

      <div class="flex justify-end mt-4 space-x-2">
        <Button variant="outline" @click="navigateTo('/application')">Cancel</Button>
        <Button @click="onSubmit">Submit</Button>
      </div>
  </div>
</template>
