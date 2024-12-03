<script setup lang="ts">
import { cn } from "@/lib/utils";
import {
  DateFormatter,
  parseDate
} from "@internationalized/date";
import { CalendarIcon } from "@radix-icons/vue";
import { toTypedSchema } from "@vee-validate/zod";
import { toDate } from "radix-vue/date";
import { useForm } from "vee-validate";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
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
import { Textarea } from "@/components/ui/textarea";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig();
const { showLoading, hideLoading } = useLoading();

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    vendor_id: z.string().min(1, "Vendor harus dipilih"),
    merek_perangkat: z.string().min(1, "Merek perangkat harus diisi"),
    nomor_nota: z.string().min(1, "Nomor nota harus diisi"),
    lokasi_penerima: z.string().min(1, "Lokasi penerima harus diisi"),
    tanggal_penerimaan: z.string().min(1, "Tanggal penerimaan harus diisi"),
    masa_garansi_mulai: z.string().min(1, "Masa garansi mulai harus diisi"),
    penanggung_jawab_perangkat: z
      .string()
      .min(1, "Penanggung jawab harus dipilih"),
    model: z.string().min(1, "Model harus diisi"),
    serial_number: z.string().min(1, "Serial number harus diisi"),
    harga_perangkat: z.string().min(1, "Harga perangkat harus diisi"),
    depresiasi_perangkat: z.string().min(1, "Depresiasi perangkat harus diisi"),
    tanda_terima: z.any(),
    kondisi_aset: z.string().min(1, "Kondisi aset harus dipilih"),
    tipe_perangkat: z.string().min(1, "Tipe perangkat harus dipilih"),
    tanggal_aktivasi_perangkat: z
      .string()
      .min(1, "Tanggal aktivasi harus diisi"),
    masa_berakhir_garansi: z
      .string()
      .min(1, "Masa berakhir garansi harus diisi"),
    hasil_pemeriksaan_perangkat: z
      .string()
      .min(1, "Hasil pemeriksaan harus diisi"),
    jangka_masa_pakai: z.string().min(1, "Jangka masa pakai harus diisi"),
    status_perangkat: z.string().min(1, "Status perangkat harus dipilih"),
    nota_pembelian: z.any(),
    detail_spesifikasi: z.string().min(1, "Detail spesifikasi harus diisi"),
    nomor_kartu_garansi: z.string().min(1, "Nomor kartu garansi harus diisi"),
    divisi_pengguna: z.string().min(1, "Divisi pengguna harus dipilih"),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});

const onSubmit = handleSubmit(async (values) => {
  showLoading();

  const formData = new FormData();
  Object.keys(values).forEach((key) => {
    formData.append(key, values[key]);
  });

  const { data, status } = await useFetch(
    config.public.API_URL + "/form-hardware",
    {
      method: "POST",
      body: formData,
    }
  );

  hideLoading();

  if (status.value == "success") {
    navigateTo("/hardware");
  }
});

const df = new DateFormatter("id-ID", {
  dateStyle: "long",
});
/* data select vendor */
const vendors = ref([]);

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

/* hold datefield values */
const receivedDate = computed({
  get: () =>
    values.tanggal_penerimaan
      ? parseDate(values.tanggal_penerimaan)
      : undefined,
  set: (val) => val,
});

const warrantyStart = computed({
  get: () =>
    values.masa_garansi_mulai
      ? parseDate(values.masa_garansi_mulai)
      : undefined,
  set: (val) => val,
});

const warrantyEnd = computed({
  get: () =>
    values.masa_berakhir_garansi
      ? parseDate(values.masa_berakhir_garansi)
      : undefined,
  set: (val) => val,
});

const activationDate = computed({
  get: () =>
    values.tanggal_aktivasi_perangkat
      ? parseDate(values.tanggal_aktivasi_perangkat)
      : undefined,
  set: (val) => val,
});
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Hardware</h1>
    <form class="grid grid-cols-2 gap-x-4 gap-y-3">
      <FormField v-slot="{ componentField }" name="vendor_id">
        <FormItem>
          <FormLabel>Vendor</FormLabel>
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
          <FormLabel>Serial Number</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="tipe_perangkat">
        <FormItem>
          <FormLabel>Tipe Perangkat</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Tipe" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <SelectItem value="SSD">SSD</SelectItem>
                <SelectItem value="Kabel Internet">Kabel Internet</SelectItem>
                <SelectItem value="RAM">RAM</SelectItem>
                <SelectItem value="HDD">HDD</SelectItem>
                <SelectItem value="Prosesor">Prosesor</SelectItem>
                <SelectItem value="Casing">Casing</SelectItem>
                <SelectItem value="Mainboard">Mainboard</SelectItem>
                <SelectItem value="FAN">FAN</SelectItem>
                <SelectItem value="Cooling Pasta">Cooling Pasta</SelectItem>
              </SelectGroup>
            </SelectContent>
          </Select>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="nomor_nota">
        <FormItem>
          <FormLabel>Nomor Nota</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="tanggal_penerimaan">
        <FormItem class="flex flex-col">
          <FormLabel>Tanggal Penerimaan</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'ps-3 text-start font-normal',
                      !receivedDate && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    receivedDate
                      ? df.format(toDate(receivedDate))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="receivedDate"
                @update:model-value="
                  (v) => setFieldValue('tanggal_penerimaan', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="lokasi_penerima">
        <FormItem>
          <FormLabel>Lokasi Penerima</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="merek_perangkat">
        <FormItem>
          <FormLabel>Merek Perangkat</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="model">
        <FormItem>
          <FormLabel>Model</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="nomor_kartu_garansi">
        <FormItem>
          <FormLabel>Nomor Kartu Garansi</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="masa_garansi_mulai">
        <FormItem class="flex flex-col">
          <FormLabel>Garansi Mulai</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'ps-3 text-start font-normal',
                      !warrantyStart && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    warrantyStart
                      ? df.format(toDate(warrantyStart))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="warrantyStart"
                @update:model-value="
                  (v) => setFieldValue('masa_garansi_mulai', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="masa_berakhir_garansi">
        <FormItem class="flex flex-col">
          <FormLabel>Garansi Berakhir</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'ps-3 text-start font-normal',
                      !warrantyEnd && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    warrantyEnd
                      ? df.format(toDate(warrantyEnd))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="warrantyEnd"
                @update:model-value="
                  (v) => setFieldValue('masa_berakhir_garansi', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="tanggal_aktivasi_perangkat">
        <FormItem class="flex flex-col">
          <FormLabel>Aktivasi Perangkat</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'ps-3 text-start font-normal',
                      !activationDate && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    activationDate
                      ? df.format(toDate(activationDate))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="activationDate"
                @update:model-value="
                  (v) =>
                    setFieldValue('tanggal_aktivasi_perangkat', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="jangka_masa_pakai">
        <FormItem>
          <FormLabel>Jangka Masa Pakai (Tahun)</FormLabel>
          <FormControl>
            <Input type="number" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="penanggung_jawab_perangkat">
        <FormItem>
          <FormLabel>Penanggung Jawab</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="status_perangkat">
        <FormItem>
          <FormLabel>Status Perangkat</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Status" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in deviceStatuses">
                  <SelectItem :value="item.value">
                    {{ item.label }}
                  </SelectItem>
                </template>
              </SelectGroup>
            </SelectContent>
          </Select>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="hasil_pemeriksaan_perangkat">
        <FormItem>
          <FormLabel>Hasil Pemeriksaan</FormLabel>
          <FormControl>
            <Textarea
              placeholder=""
              class="resize-none"
              v-bind="componentField"
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="detail_spesifikasi">
        <FormItem>
          <FormLabel>Detail Spesifikasi</FormLabel>
          <FormControl>
            <Textarea
              placeholder=""
              class="resize-none"
              v-bind="componentField"
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="harga_perangkat">
        <FormItem>
          <FormLabel>Harga Perangkat</FormLabel>
          <FormControl>
            <Input type="number" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="depresiasi_perangkat">
        <FormItem>
          <FormLabel>Depresiasi Perangkat</FormLabel>
          <FormControl>
            <Input type="number" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="divisi_pengguna">
        <FormItem>
          <FormLabel>Divisi Pengguna</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Divisi" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in divisions">
                  <SelectItem :value="item.value">
                    {{ item.label }}
                  </SelectItem>
                </template>
              </SelectGroup>
            </SelectContent>
          </Select>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="kondisi_aset">
        <FormItem>
          <FormLabel>Kondisi Aset</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Kondisi" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in assetConditions">
                  <SelectItem :value="item.value">
                    {{ item.label }}
                  </SelectItem>
                </template>
              </SelectGroup>
            </SelectContent>
          </Select>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="tanda_terima">
        <FormItem>
          <FormLabel>Tanda Terima</FormLabel>
          <FormControl>
            <Input
              type="file"
              accept=".pdf,.jpg,.jpeg,.png"
              @change="(e) => setFieldValue('tanda_terima', e.target.files[0])"
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="nota_pembelian">
        <FormItem>
          <FormLabel>Nota Pembelian</FormLabel>
          <FormControl>
            <Input
              type="file"
              accept=".pdf,.jpg,.jpeg,.png"
              @change="
                (e) => setFieldValue('nota_pembelian', e.target.files[0])
              "
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
    </form>

    <div class="flex justify-end mt-4 space-x-2">
      <Button @click="onSubmit">Submit</Button>
    </div>
  </div>
</template>
