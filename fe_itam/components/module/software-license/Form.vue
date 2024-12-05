<script setup lang="ts">
import { useToast } from "@/components/ui/toast/use-toast";
import { cn } from "@/lib/utils";
import {
  DateFormatter,
  parseDate
} from "@internationalized/date";
import { CalendarIcon } from "@radix-icons/vue";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { toDate } from "radix-vue/date";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig();
const route = useRoute();
const { showLoading, hideLoading } = useLoading();
const { toast } = useToast();
const vendors = ref([]);
let exData = <any>{}

const df = new DateFormatter("id-ID", {
  dateStyle: "long",
});

/* hold datefield value */
const waktuPembelian = computed({
  get: () =>
    values.waktu_pembelian ? parseDate(values.waktu_pembelian) : undefined,
  set: (val) => val,
});

const waktuAktivasi = computed({
  get: () =>
    values.waktu_aktivasi ? parseDate(values.waktu_aktivasi) : undefined,
  set: (val) => val,
});

const tanggalExpired = computed({
  get: () =>
    values.tanggal_expired ? parseDate(values.tanggal_expired) : undefined,
  set: (val) => val,
});

/* data select vendor */

const getVendorData = async () => {
  try {
    const { data, status } = await useFetch(config.public.API_URL + "/vendor");

    if (status.value == "success" && data.value?.data?.length) {
      vendors.value = data.value.data.map((item) => {
        return {
          value: item.id,
          label: item.nama_pic,
        };
      });
    }
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
};

onMounted(() => {
  const retryInterval = setInterval(() => {
    if (!vendors.value?.length) getVendorData();
    else clearInterval(retryInterval);
  }, 500);
});

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    vendor_id: z.number().min(1),
    serial_number: z.string().min(1),
    merk: z.string().min(1),
    model: z.string().min(1),
    nomor_nota: z.string().min(1),
    waktu_pembelian: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), {
        message: "Invalid date format, expected YYYY-MM-DD",
      })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    SN_perangkat_terpasang: z.string().min(1),
    waktu_aktivasi: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), {
        message: "Invalid date format, expected YYYY-MM-DD",
      })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tanggal_expired: z
      .string()
      .refine((val) => /^\d{4}-\d{2}-\d{2}$/.test(val), {
        message: "Invalid date format, expected YYYY-MM-DD",
      })
      .transform((val) => new Date(`${val}T00:00:00Z`).toISOString()),
    tipe_kepemilikan_aset: z.string().min(1),
    kategori_lisensi: z.string().min(1),
    versi_lisensi: z.string().min(1),
    maksimal_user_aplikasi: z.number().min(1),
    maksimal_perangkat_lisensi: z.number().min(1),
    tipe_lisensi: z.string().min(1),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
  initialValues: {
    vendor_id: undefined,
    serial_number: "",
    merk: "",
    model: "",
    nomor_nota: "",
    waktu_pembelian: undefined,
    SN_perangkat_terpasang: "",
    waktu_aktivasi: undefined,
    tanggal_expired: undefined,
    tipe_kepemilikan_aset: "",
    kategori_lisensi: "",
    versi_lisensi: "",
    maksimal_user_aplikasi: undefined,
    maksimal_perangkat_lisensi: undefined,
    tipe_lisensi: "",
  }
});

const dataId = route.params.id;
const endpoint = `/asset-lisensi${props.type == "new" ? "" : "/" + dataId}`;

const onSubmit = handleSubmit(async (values) => {
  showLoading();

  try {
    const { data, status } = await useFetch(config.public.API_URL + endpoint, {
      method: props.type == "new" ? "POST" : "PATCH",
      body: values,
    });

    if (status.value == "success") {
      toast({
        title: "Success",
        description: `Data submitted successfully`,
      });
      navigateTo("/application");
    } else {
      toast({
        title: "Failed",
        description: `Error when submitting data`,
      });
    }
  } catch (error) {
    console.error("Error occured:", error);
  }

  hideLoading();
});

const getExistingData = async () => {
  showLoading()

  try {
    // First, ensure vendors are loaded
    await getVendorData()

    const { data, status } = await useFetch(config.public.API_URL + endpoint);

    if (status.value == 'success' && data.value?.data) {
      exData = data.value.data

      // Ensure dates are formatted correctly
      const formatDate = (dateString: string) => {
        if (!dateString) return undefined;
        // Parse the date and convert to ISO string in YYYY-MM-DD format
        const date = new Date(dateString);
        return date.toISOString().split('T')[0];
      };

      // Set vendor only if it exists in the vendors list
      const vendorExists = vendors.value.some(v => v.value === exData.asset.vendor_id)
      if (vendorExists) {
        setFieldValue('vendor_id', exData.asset.vendor_id)
      } else {
        console.warn('Vendor not found in the list')
      }

      setFieldValue('serial_number', exData.asset.serial_number)
      setFieldValue('merk', exData.asset.merk)
      setFieldValue('model', exData.asset.model)
      setFieldValue('nomor_nota', exData.asset.nomor_nota)

      // Format dates consistently
      setFieldValue('waktu_pembelian', formatDate(exData.waktu_pembelian))
      setFieldValue('SN_perangkat_terpasang', exData.SN_perangkat_terpasang)
      setFieldValue('waktu_aktivasi', formatDate(exData.waktu_aktivasi))
      setFieldValue('tanggal_expired', formatDate(exData.tanggal_expired))

      setFieldValue('tipe_kepemilikan_aset', exData.tipe_kepemilikan_aset)
      setFieldValue('kategori_lisensi', exData.kategori_lisensi)
      setFieldValue('versi_lisensi', exData.versi_lisensi)
      setFieldValue('maksimal_user_aplikasi', exData.maksimal_user_aplikasi)
      setFieldValue('maksimal_perangkat_lisensi', exData.maksimal_perangkat_lisensi)
      setFieldValue('tipe_lisensi', exData.tipe_lisensi)
    }
  } catch (error) {
    console.error("Error occurred:", error);
    toast({
      title: "Error",
      description: "Failed to load existing data"
    });
  }

  hideLoading()
}

// Modify the existing code to ensure getVendorData is called before getExistingData
if (props.type == 'edit') {
  // Ensure vendors are loaded first
  onMounted(async () => {
    await getVendorData();
    getExistingData();
  });
}
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">
      {{ props.type == "new" ? "Registrasi" : "Edit" }} Lisensi Software
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

        <FormField v-slot="{ componentField }" name="serial_number">
          <FormItem>
            <FormLabel>Nomor Seri</FormLabel>
            <FormControl>
              <Input type="text" id="serial_number" v-bind="componentField" required
                placeholder="Masukkan Nomor Seri" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="merk">
          <FormItem>
            <FormLabel>Merk</FormLabel>
            <FormControl>
              <Input type="text" id="merk" v-bind="componentField" required placeholder="Masukkan Merek" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="model">
          <FormItem>
            <FormLabel>Model</FormLabel>
            <FormControl>
              <Input type="text" id="model" v-bind="componentField" required placeholder="Masukkan Model" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_nota">
          <FormItem>
            <FormLabel>Nomor Nota</FormLabel>
            <FormControl>
              <Input type="text" id="nomor_nota" v-bind="componentField" required placeholder="Masukkan Nomor Nota" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="waktu_pembelian">
          <FormItem class="flex flex-col">
            <FormLabel>Waktu Pembelian</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !waktuPembelian && 'text-muted-foreground'
                  )
                    ">
                    <span>{{
                      waktuPembelian
                        ? df.format(toDate(waktuPembelian))
                        : "Pilih tanggal"
                    }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden />
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="waktuPembelian" calendar-label="Waktu Pembelian" @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('waktu_pembelian', v.toString());
                    } else {
                      setFieldValue('waktu_pembelian', undefined);
                    }
                  }
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="SN_perangkat_terpasang">
          <FormItem>
            <FormLabel>SN Perangkat Terpasang</FormLabel>
            <FormControl>
              <Input type="text" id="SN_perangkat_terpasang" v-bind="componentField" required
                placeholder="Masukkan Nomor Seri Terpasang" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="waktu_aktivasi">
          <FormItem class="flex flex-col">
            <FormLabel>Waktu Aktivasi</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !waktuAktivasi && 'text-muted-foreground'
                  )
                    ">
                    <span>{{
                      waktuAktivasi
                        ? df.format(toDate(waktuAktivasi))
                        : "Pilih tanggal"
                    }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden />
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="waktuAktivasi" calendar-label="Waktu Aktivasi" @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('waktu_aktivasi', v.toString());
                    } else {
                      setFieldValue('waktu_aktivasi', undefined);
                    }
                  }
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_expired">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Expired</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !tanggalExpired && 'text-muted-foreground'
                  )
                    ">
                    <span>{{
                      tanggalExpired
                        ? df.format(toDate(tanggalExpired))
                        : "Pilih tanggal"
                    }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                  <input hidden />
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="tanggalExpired" calendar-label="Tanggal Expired" @update:model-value="(v) => {
                    if (v) {
                      setFieldValue('tanggal_expired', v.toString());
                    } else {
                      setFieldValue('tanggal_expired', undefined);
                    }
                  }
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipe_kepemilikan_aset">
          <FormItem>
            <FormLabel>Tipe Kepemilikan Aset</FormLabel>
            <FormControl>
              <Input type="text" id="tipe_kepemilikan_aset" v-bind="componentField" required
                placeholder="Masukkan Tipe Kepemilikan Aset" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="kategori_lisensi">
          <FormItem>
            <FormLabel>Kategori Lisensi</FormLabel>
            <FormControl>
              <Input type="text" id="kategori_lisensi" v-bind="componentField" required
                placeholder="Masukkan Kategori Lisensi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="versi_lisensi">
          <FormItem>
            <FormLabel>Versi Lisensi</FormLabel>
            <FormControl>
              <Input type="text" id="versi_lisensi" v-bind="componentField" required
                placeholder="Masukkan Versi Lisensi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="maksimal_user_aplikasi">
          <FormItem>
            <FormLabel>Maksimal User Aplikasi</FormLabel>
            <FormControl>
              <Input type="number" id="maksimal_user_aplikasi" v-bind="componentField" required
                placeholder="Masukkan Maksimal User Aplikasi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="maksimal_perangkat_lisensi">
          <FormItem>
            <FormLabel>Maksimal Perangkat Lisensi</FormLabel>
            <FormControl>
              <Input type="number" id="maksimal_perangkat_lisensi" v-bind="componentField" required
                placeholder="Masukkan Maksimal Perangkat Lisensi" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipe_lisensi">
          <FormItem>
            <FormLabel>Tipe Lisensi</FormLabel>
            <FormControl>
              <Input type="text" id="tipe_lisensi" v-bind="componentField" required
                placeholder="Masukkan Tipe Lisensi" />
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
