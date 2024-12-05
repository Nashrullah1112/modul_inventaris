<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { CalendarIcon } from "@radix-icons/vue";
import { cn } from "@/lib/utils";
import {
  DateFormatter,
  parseDate
} from "@internationalized/date";
import { toDate } from "radix-vue/date";
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
import { toTypedSchema } from "@vee-validate/zod";
import { useToast } from "~/components/ui/toast";

// Props definition
const props = defineProps<{
  type: string;
  data: any[];
}>();
const route = useRoute();
const config = useRuntimeConfig();
const { showLoading, hideLoading } = useLoading();
const { toast } = useToast();
const assetSupplier = ref<Supplier[]>([]);
const assetUser = ref<User[]>([]);
const assetDivision = ref<Division[]>([]);


interface User {
  id: number;
  nama: string;
}

interface Division {
  id: number;
  nama: string;
}

interface Supplier {
  id: number
  nama_pic: string
}

const df = new DateFormatter("id-ID", {
  dateStyle: "long",
});


// Form validation schema with improved field names
const formSchema = toTypedSchema(
  z.object({
    nama_supplier: z.string().min(1, "Supplier harus dipilih"),
    nomor_seri: z.string().min(1, "Nomor seri harus diisi"),
    merek: z.string().min(1, "Merek harus diisi"),
    model: z.string().min(1, "Model harus diisi"),
    nomor_nota: z.string().min(1, "Nomor nota harus diisi"),
    lokasi_penerima: z.string().min(1, "Lokasi penerima harus diisi"),
    waktu_penerimaan: z.string().min(1, "Tanggal penerimaan harus diisi"),
    dokumen_terima: z.any().nullable(),
    tipe_aset: z.string().min(1, "Tipe aset harus dipilih"),
    tanggal_aktivasi_aset: z.string().min(1, "Tanggal aktivasi harus diisi"),
    hasil_pemeriksaan_aset: z.string().min(1, "Hasil pemeriksaan harus diisi"),
    masa_garansi_mulai: z.string().min(1, "Masa garansi mulai harus diisi"),
    nomor_kartu_garansi: z.string().min(1, "Nomor kartu garansi harus diisi"),
    prosesor: z.string().min(1, "Prosesor harus diisi"),
    kapasitas_ram: z.number().min(0, "Kapasitas RAM harus diisi"),
    kapasitas_rom: z.number().min(0, "Kapasitas ROM harus diisi"),
    tipe_ram: z.string().min(1, "Tipe RAM harus dipilih"),
    tipe_penyimpanan: z.string().min(1, "Tipe penyimpanan harus dipilih"),
    status_aset: z.string().min(1, "Status aset harus dipilih"),
    nilai_aset: z.number().min(0, "Nilai aset harus diisi"),
    nilai_sisa: z.number().min(0, "Nilai sisa aset harus diisi"),
    jangka_masa_pakai: z.number().min(1, "Jangka masa pakai harus dipilih"),
    waktu_aset_keluar: z.string().min(1, "Waktu aset keluar harus diisi"),
    kondisi_aset_keluar: z.string().min(1, "Kondisi aset keluar harus dipilih"),
    nota_pembelian: z.any().nullable(),
    divisi_pengguna: z.string().min(1, "Divisi harus dipilih"),
    penanggung_jawab_aset: z.string().min(1, "Penanggung jawab harus dipilih"),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});

const assetType = [
  {
    id: "1",
    nama: "Laptop",
  },
  {
    id: "2",
    nama: "Smartphone",
  },
  {
    id: "3",
    nama: "Tablet",
  }
]

const assetUsagePeriod = [
  {
    id: "1",
    nama: "1 Tahun",
  },
  {
    id: "2",
    nama: "2 Tahun",
  },
  {
    id: "3",
    nama: "3 Tahun",
  },
  {
    id: "4",
    nama: "4 Tahun",
  },
  {
    id: "5",
    nama: "5 Tahun",
  },
  {
    id: "6",
    nama: "6 Tahun",
  },
  {
    id: "7",
    nama: "7 Tahun",
  },
  {
    id: "8",
    nama: "8 Tahun",
  },
  {
    id: "9",
    nama: "9 Tahun",
  },
  {
    id: "10",
    nama: "10 Tahun",
  }
];

const assetCondition = ["Baik", "Rusak", "Perlu Perbaikan"];
const assetStatus = ["Aktif", "Tidak Aktif", "Sedang Digunakan"];
const ramCapacity = ["2 GB", "4 GB", "8 GB", "16 GB", "32 GB", "64 GB", "128 GB", "256 GB"];
const romCapacity = ["16 GB", "32 GB", "64 GB", "128 GB", "256 GB", "512 GB", "1 TB", "2 TB", "4 TB", "8 TB"];
const ramType = ["DDR3", "DDR4", "DDR5"];
const storageType = ["HDD", "SSD"];

// Initialize form
const form = useForm({
  validationSchema: formSchema,
  initialValues: {
    nama_supplier: "",
    nomor_seri: "",
    merek: "",
    model: "",
    nomor_nota: "",
    lokasi_penerima: "",
    waktu_penerimaan: "",
    dokumen_terima: null,
    tipe_aset: "",
    tanggal_aktivasi_aset: "",
    hasil_pemeriksaan_aset: "",
    masa_garansi_mulai: "",
    nomor_kartu_garansi: "",
    prosesor: "",
    kapasitas_ram: 0,
    kapasitas_rom: 0,
    tipe_ram: "",
    tipe_penyimpanan: "",
    status_aset: "",
    nilai_aset: 0,
    nilai_sisa: 0,
    jangka_masa_pakai: 0,
    waktu_aset_keluar: "",
    kondisi_aset_keluar: "",
    nota_pembelian: null,
    divisi_pengguna: "",
    penanggung_jawab_aset: ""
  },
});


const fetchSupplierData = async () => {
  try {
    const response = await $fetch<{ data: Supplier[] }>(config.public.API_URL + '/vendor', {
      method: 'GET'
    });

    if (response && response.data) {
      assetSupplier.value = response.data;
    }
  } catch (error) {
    console.error('Error occurred:', error);
  }
}

const fetchUserData = async () => {
  try {
    const response = await $fetch<{ data: User[] }>(config.public.API_URL + '/user', {
      method: 'GET'
    });

    if (response && response.data) {
      assetUser.value = response.data;
    }
  } catch (error) {
    console.error('Error occurred:', error);
  }
}

const fetchDivisionData = async () => {
  try {
    const response = await $fetch<{ data: Division[] }>(config.public.API_URL + '/divisi', {
      method: 'GET'
    });

    if (response && response.data) {
      assetDivision.value = response.data;
    }
  } catch (error) {
    console.error('Error occurred:', error);
  }
}


onMounted(async () => {
  showLoading();
  await fetchSupplierData();
  await fetchUserData();
  await fetchDivisionData();
  hideLoading();
});
// Form submission handler
const onSubmit = handleSubmit(async (values: any) => {
  console.log("Form submission triggered"); // Debug log

  try {
    // Log the entire values object to see what's being captured
    console.log("Raw submitted values:", values);

    // Your existing detailed logging
    console.log("Submitted Asset Registration Data:", {
      supplierInfo: {
        supplierName: values.nama_supplier,
        serialNumber: values.nomor_seri,
        brand: values.merek,
        model: values.model,
        invoiceNumber: values.nomor_nota,
        receiverLocation: values.lokasi_penerima,
        receivedDate: values.waktu_penerimaan,
      },
      // ... rest of the existing logging
    });

    // Prepare form data for API submission
    const formData = new FormData();

    Object.keys(values).forEach(key => {
      if (values[key] instanceof File) {
        formData.append(key, values[key]);
      } else {
        // Convert to string, handling potential undefined/null
        formData.append(key, values[key]?.toString() || '');
      }
    });

    console.log("FormData entries:", Object.fromEntries(formData.entries()));

    // Optional: Uncomment if you want to actually submit
    // const response = await axios.post('/api/assets', formData, {
    //   headers: { 'Content-Type': 'multipart/form-data' }
    // });

    console.log("Form submitted successfully");

  } catch (error) {
    console.error("Detailed error in form submission:", error);
  }
});


const receivedDate = computed({
  get: () =>
    values.waktu_penerimaan
      ? parseDate(values.waktu_penerimaan)
      : undefined,
  set: (val) => val,
});

const activationDate = computed({
  get: () =>
    values.tanggal_aktivasi_aset
      ? parseDate(values.tanggal_aktivasi_aset)
      : undefined,
  set: (val) => val,
});

const warrantyStartDate = computed({
  get: () =>
    values.masa_garansi_mulai
      ? parseDate(values.masa_garansi_mulai)
      : undefined,
  set: (val) => val,
});

const assetExitDate = computed({
  get: () =>
    values.waktu_aset_keluar
      ? parseDate(values.waktu_aset_keluar)
      : undefined,
  set: (val) => val,
});
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">
      Registrasi Aset Elektronik
    </h1>

    <Form @submit="onSubmit()" :form="form" class="space-y-6">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-4">
        <FormField v-slot="{ componentField }" name="nama_supplier">
          <FormItem>
            <FormLabel>Nama Supplier</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Nama Supplier" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetSupplier" :key="item.id">
                    <SelectItem :value="item.id.toString()">
                      {{ item.nama_pic }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_seri">
          <FormItem>
            <FormLabel>Nomor Seri</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Nomor Seri" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="merek">
          <FormItem>
            <FormLabel>Merek</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Merek" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="model">
          <FormItem>
            <FormLabel>Model</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Model" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_nota">
          <FormItem>
            <FormLabel>Nomor Nota</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Nomor Nota" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="lokasi_penerima">
          <FormItem>
            <FormLabel>Lokasi Penerima</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Lokasi Penerima" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="waktu_penerimaan">
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
                <Calendar v-model="receivedDate" @update:model-value="(v) => setFieldValue('waktu_penerimaan', v?.toString())
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipe_aset">
          <FormItem>
            <FormLabel>Tipe Aset</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Tipe Aset" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetType" :key="item">
                    <SelectItem :value="item.nama">
                      {{ item.nama }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField name="tanggal_aktivasi_aset">
          <FormItem class="flex flex-col">
            <FormLabel>Tanggal Aktivasi Aset</FormLabel>
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
                <Calendar v-model="activationDate" @update:model-value="(v) => setFieldValue('tanggal_aktivasi_aset', v?.toString())
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>


        <FormField v-slot="{ componentField }" name="hasil_pemeriksaan_aset">
          <FormItem>
            <FormLabel>Hasil Pemeriksaan Aset</FormLabel>
            <FormControl>
              <Textarea placeholder="Masukkan Hasil Pemeriksaan" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="tanggal_aktivasi_aset">
          <FormItem class="flex flex-col">
            <FormLabel>Masa Garansi Mulai</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !warrantyStartDate && 'text-muted-foreground'
                  )
                    ">
                    <span>{{
                      warrantyStartDate
                        ? df.format(toDate(warrantyStartDate))
                        : "Pilih Tanggal"
                    }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="warrantyStartDate" @update:model-value="(v) => setFieldValue('masa_garansi_mulai', v?.toString())
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_kartu_garansi">
          <FormItem>
            <FormLabel>Nomor Kartu Garansi</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Nomor Kartu Garansi" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="prosesor">
          <FormItem>
            <FormLabel>Prosesor</FormLabel>
            <FormControl>
              <Input type="text" placeholder="Masukkan Prosesor" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="kapasitas_ram">
          <FormItem>
            <FormLabel>Kapasitas RAM</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Kapasitas RAM" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in ramCapacity" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="kapasitas_rom">
          <FormItem>
            <FormLabel>Kapasitas ROM</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Kapasitas ROM" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in romCapacity" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipe_ram">
          <FormItem>
            <FormLabel>Tipe RAM</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Tipe RAM" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in ramType" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="tipe_penyimpanan">
          <FormItem>
            <FormLabel>Tipe Penyimpanan</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Tipe Penyimpanan" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in storageType" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="status_aset">
          <FormItem>
            <FormLabel>Status Aset</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Status Aset" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetStatus" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nilai_aset">
          <FormItem>
            <FormLabel>Nilai Aset</FormLabel>
            <FormControl>
              <Input type="number" placeholder="Masukkan Nilai Aset" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nilai_sisa">
          <FormItem>
            <FormLabel>Nilai Sisa Aset</FormLabel>
            <FormControl>
              <Input type="number" placeholder="Masukkan Nilai Sisa Aset" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="jangka_masa_pakai">
          <FormItem>
            <FormLabel>Jangka Masa Pakai</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Jangka Masa Pakai" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetUsagePeriod" :key="item">
                    <SelectItem :value="item.nama">
                      {{ item.nama }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField name="waktu_aset_keluar">
          <FormItem class="flex flex-col">
            <FormLabel>Waktu Aset Keluar</FormLabel>
            <Popover>
              <PopoverTrigger as-child>
                <FormControl>
                  <Button variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !assetExitDate && 'text-muted-foreground'
                  )
                    ">
                    <span>{{
                      assetExitDate
                        ? df.format(toDate(assetExitDate))
                        : "Pilih Tanggal"
                    }}</span>
                    <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                  </Button>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent class="w-auto p-0">
                <Calendar v-model="assetExitDate" @update:model-value="(v) => setFieldValue('waktu_aset_keluar', v?.toString())
                  " />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="kondisi_aset_keluar">
          <FormItem>
            <FormLabel>Kondisi Aset Keluar</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Kondisi Aset Keluar" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetCondition" :key="item">
                    <SelectItem :value="item">
                      {{ item }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="divisi_pengguna">
          <FormItem>
            <FormLabel>Divisi Pengguna</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Divisi Pengguna" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetDivision" :key="item">
                    <SelectItem :value="item.id.toString()">
                      {{ item.nama }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="penanggung_jawab_aset">
          <FormItem>
            <FormLabel>Penanggung Jawab Aset</FormLabel>
            <Select v-bind="componentField">
              <SelectTrigger>
                <SelectValue placeholder="Pilih Penanggung Jawab Aset" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  <template v-for="item in assetUser" :key="item">
                    <SelectItem :value="item.id.toString()">
                      {{ item.nama }}
                    </SelectItem>
                  </template>
                </SelectGroup>
              </SelectContent>
            </Select>
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="dokumen_terima">
          <FormItem>
            <FormLabel>Dokumen Terima</FormLabel>
            <FormControl>
              <Input type="file" accept=".pdf,.jpg,.jpeg,.png" @change="(e: any) => {
                const file = e.target.files[0];
                setFieldValue('dokumen_terima', file ? file : null);
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

      </div>

      <div class="flex justify-end space-x-4 mt-6">
        <Button variant="outline" type="button" class="bg-gray-100 hover:bg-gray-200">
          Cancel
        </Button>
        <Button type="submit" class="text-white hover:bg-blue-700 transition-colors duration-300">
          Submit
        </Button>
      </div>
    </Form>
  </div>
</template>

<style scoped>
/* Custom styles for enhanced aesthetics */
</style>