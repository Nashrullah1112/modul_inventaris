<script setup lang="ts">
import { ref, onMounted } from "vue";
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
import { useToast } from "~/components/ui/toast";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig();
const router = useRouter();
const { showLoading, hideLoading } = useLoading();
const { toast } = useToast();
const vendors = ref<Vendor[]>([]);
const divisions = ref<Division[]>([]);
const users = ref<User[]>([]);
const assetConditions = [
  { label: "Baik", value: "Baik" },
  { label: "Rusak", value: "Rusak" },
  { label: "Hilang", value: "Hilang" },
];

const deviceStatuses = [
  { label: "Aktif", value: "Aktif" },
  { label: "Tidak Aktif", value: "Tidak Aktif" },
];

interface User {
  id: number;
  nama: string;
}

interface Division {
  id: number;
  nama: string;
}

interface Vendor {
  id: number;
  nama: string;
}


/* handle form */
const formSchema = toTypedSchema(
  z.object({
    vendor_id: z.number().min(1, "Vendor harus dipilih"),
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
    harga_perangkat: z.number().min(1, "Harga perangkat harus diisi"),
    depresiasi_perangkat: z.number().min(1, "Depresiasi perangkat harus diisi"),
    tanda_terima: z.any(),
    kondisi_aset: z.string().min(1, "Kondisi aset harus dipilih"),
    tipe_perangkat: z.string().min(1, "Tipe perangkat harus dipilih"),
    tanggal_aset_keluar: z.string().date().min(1, "Tanggal aset keluar harus diisi"),
    tanggal_aktivasi_perangkat: z
      .string()
      .min(1, "Tanggal aktivasi harus diisi"),
    masa_berakhir_garansi: z
      .string()
      .min(1, "Masa berakhir garansi harus diisi"),
    hasil_pemeriksaan_perangkat: z
      .string()
      .min(1, "Hasil pemeriksaan harus diisi"),
    jangka_masa_pakai: z.number().min(1, "Jangka masa pakai harus diisi"),
    status_perangkat: z.string().min(1, "Status perangkat harus dipilih"),
    nota_pembelian: z.any(),
    detail_spesifikasi: z.string().min(1, "Detail spesifikasi harus diisi"),
    nomor_kartu_garansi: z.string().min(1, "Nomor kartu garansi harus diisi"),
    divisi_pengguna: z.number().min(1, "Divisi pengguna harus dipilih"),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});


const endpoint = props.type === "new" ? "/form-hardware" : "/asset/hardware" + props.data.id;

``
const onSubmit = handleSubmit(async (values) => {
  try {
    showLoading();

    const formData = new FormData();

    (Object.keys(values) as Array<keyof typeof values>).forEach(key => {
      const value = values[key];

      if ((key === 'tanda_terima' || key === 'nota_pembelian') && value instanceof File) {
        formData.append(key, value);
      } else if (value !== undefined && value !== null) {
        formData.append(key, String(value));
      }
    });

    const { data, status, error } = await useFetch(
      config.public.API_URL + endpoint,
      {
        method: props.type == "new" ? "POST" : "PATCH",
        body: formData,
      }
    )

    if (status.value === 'success') {
      toast({
        title: 'Success',
        description: `Data submitted successfully`,
      })
      router.push('/hardware');
    }

  } catch (error) {
    hideLoading();
    toast({
      title: 'Gagal Menyimpan Data',
      description: `Gagal menyimpan data, silahkan coba lagi`,
      variant: 'destructive'
    })
    console.log("Terjadi kesalahan:", error);
  } finally {
    hideLoading();
  }
});

const df = new DateFormatter("id-ID", {
  dateStyle: "long",
});

async function fetchDropdownData() {
  try {
    showLoading();

    const userResponse = await $fetch<ApiResponse<{ id: number; nama: string }>>(
      config.public.API_URL + "/user"
    );

    users.value = userResponse.data.map(user => ({
      id: user.id,
      nama: user.nama
    }));

    const divisionResponse = await $fetch<ApiResponse<{ id: number; nama: string }>>(
      config.public.API_URL + "/divisi"
    );

    divisions.value = divisionResponse.data.map(division => ({
      id: division.id,
      nama: division.nama
    }));

    const vendorResponse = await $fetch<ApiResponse<{ id: number; nama_pic: string }>>(
      config.public.API_URL + "/vendor"
    );

    vendors.value = vendorResponse.data.map(vendor => ({
      id: vendor.id,
      nama: vendor.nama_pic
    }));
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  } finally {
    hideLoading();
  }
}


onMounted(() => {
  fetchDropdownData();
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

const assetOutDate = computed({
  get: () =>
    values.tanggal_aset_keluar
      ? parseDate(values.tanggal_aset_keluar)
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
                  <template v-for="item in vendors" :key="item.id">
                    <SelectItem :value="item.id">
                      {{ item.nama }}
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
                <Button variant="outline" :class="cn(
                  'ps-3 text-start font-normal',
                  !receivedDate && 'text-muted-foreground'
                )
                  ">
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
              <Calendar v-model="receivedDate" @update:model-value="(v) => setFieldValue('tanggal_penerimaan', v?.toString())
                " />
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
                <Button variant="outline" :class="cn(
                  'ps-3 text-start font-normal',
                  !warrantyStart && 'text-muted-foreground'
                )
                  ">
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
              <Calendar v-model="warrantyStart" @update:model-value="(v) => setFieldValue('masa_garansi_mulai', v?.toString())
                " />
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
                <Button variant="outline" :class="cn(
                  'ps-3 text-start font-normal',
                  !warrantyEnd && 'text-muted-foreground'
                )
                  ">
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
              <Calendar v-model="warrantyEnd" @update:model-value="(v) => setFieldValue('masa_berakhir_garansi', v?.toString())
                " />
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
                <Button variant="outline" :class="cn(
                  'ps-3 text-start font-normal',
                  !activationDate && 'text-muted-foreground'
                )
                  ">
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
              <Calendar v-model="activationDate" @update:model-value="(v) =>
                setFieldValue('tanggal_aktivasi_perangkat', v?.toString())
                " />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="tanggal_aset_keluar">
        <FormItem class="flex flex-col">
          <FormLabel>Tanggal Aset Keluar</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button variant="outline" :class="cn(
                  'ps-3 text-start font-normal',
                  !assetOutDate && 'text-muted-foreground'
                )
                  ">
                  <span>{{
                    assetOutDate
                      ? df.format(toDate(assetOutDate))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar v-model="assetOutDate" @update:model-value="(v) =>
                setFieldValue('tanggal_aset_keluar', v?.toString())
                " />
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
          <FormLabel>Penanggung Jawab Perangkat</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Penanggung Jawab Perangkat" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in users">
                  <SelectItem :value="item.nama">
                    {{ item.nama }}
                  </SelectItem>
                </template>
              </SelectGroup>
            </SelectContent>
          </Select>
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
            <Textarea placeholder="" class="resize-none" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="detail_spesifikasi">
        <FormItem>
          <FormLabel>Detail Spesifikasi</FormLabel>
          <FormControl>
            <Textarea placeholder="" class="resize-none" v-bind="componentField" />
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
              <SelectGroup label="Divisions">
                <SelectItem v-for="item in divisions" :key="item.id" :value="item.id">
                  {{ item.nama }}
                </SelectItem>
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
            <Input type="file" accept=".pdf,.jpg,.jpeg,.png" @change="(e: any) => {
              const file = e.target.files[0];
              setFieldValue('tanda_terima', file ? file : null);
            }" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="nota_pembelian">
        <FormItem>
          <FormLabel>Nota Pembelian</FormLabel>
          <FormControl>
            <Input type="file" accept=".pdf,.jpg,.jpeg,.png" @change="(e: any) => {
              const file = e.target.files[0];
              setFieldValue('nota_pembelian', file ? file : null);
            }" />
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
