<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { toast } from "@/components/ui/toast";
import { useForm } from "vee-validate";
import { computed, ref } from 'vue';
import * as z from "zod";
import { useRouter } from 'vue-router';

const router = useRouter();


const props = defineProps<{
  type: string;
  data: any[];
}>();

const emit = defineEmits(['asset-registered']);

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

const currentStep = ref(1);

const step1Fields = [
  'namasupllier',
  'tipeasset',
  'notaNumber',
  'kondisiasset',
  'status',
  'lokasipenerima',
  'nilaiasset',
  'achievedDate',
  'tandaterima',
  'hasilpemeriksaan'
];

const step2Fields = [
  'tanggalaktivasi',
  'penanggungjawabasset',
  'divpengguna',
  'lokasipenyimpanan',
  'jangkamasapakaiaset',
  'model',
  'tanggalasetkeluar',
  'nomorgaransi',
  'merk',
  'garansimulai',
  'garansiakhir',
  'serialnumber',
  'prosesor',
  'ram',
  'tiperam',
  'rom',
  'notapembelian'
];

const getFieldLabel = (field: string) => {
  const labels = {
    namasupllier: 'Nama Supplier',
    tipeasset: 'Tipe Aset',
    notaNumber: 'Nomor Nota',
    kondisiasset: 'Kondisi Aset',
    status: 'Status',
    lokasipenerima: 'Lokasi Penerima',
    nilaiasset: 'Nilai Aset',
    achievedDate: 'Tanggal Penerimaan',
    tandaterima: 'Tanda Terima',
    hasilpemeriksaan: 'Hasil Pemeriksaan',
    tanggalaktivasi: 'Tanggal Aktivasi',
    penanggungjawabasset: 'Penanggung Jawab Aset',
    divpengguna: 'Divisi Pengguna',
    lokasipenyimpanan: 'Lokasi Penyimpanan',
    jangkamasapakaiaset: 'Jangka Masa Pakai Aset',
    model: 'Model',
    tanggalasetkeluar: 'Tanggal Aset Keluar',
    nomorgaransi: 'Nomor Garansi',
    merk: 'Merek',
    garansimulai: 'Garansi Mulai',
    garansiakhir: 'Garansi Akhir',
    serialnumber: 'Serial Number',
    prosesor: 'Prosesor',
    ram: 'RAM',
    tiperam: 'Tipe RAM',
    rom: 'ROM',
    notapembelian: 'Nota Pembelian'
  };
  return labels[field] || field;
};

const getOptionsForField = (field: string) => {
  const options = {
    namasupllier: [
      { value: 'supplier1', label: 'Supplier 1' },
      { value: 'supplier2', label: 'Supplier 2' },
    ],
    tipeasset: [
      { value: 'elektronik', label: 'Elektronik' },
      { value: 'komputer', label: 'Komputer' },
    ],
    kondisiasset: [
      { value: 'baru', label: 'Baru' },
      { value: 'bekas', label: 'Bekas' },
    ],
    status: [
      { value: 'aktif', label: 'Aktif' },
      { value: 'nonaktif', label: 'Non-Aktif' },
    ],
    divpengguna: [
      { value: 'it', label: 'IT' },
      { value: 'marketing', label: 'Marketing' },
    ],
    jangkamasapakaiaset: [
      { value: '1tahun', label: '1 Tahun' },
      { value: '2tahun', label: '2 Tahun' },
    ],
    penanggungjawabasset: [
      { value: 'manager1', label: 'Manager 1' },
      { value: 'manager2', label: 'Manager 2' },
    ],
    tiperam: [
      { value: 'ddr3', label: 'DDR3' },
      { value: 'ddr4', label: 'DDR4' },
      { value: 'ddr5', label: 'DDR5' },
    ]
  };
  return options[field] || [];
};

const onSubmit = async (values) => {
  try {

    await new Promise(resolve => setTimeout(resolve, 1000));


    emit('asset-registered', values);


    toast({
      title: "Registrasi Aset Berhasil",
      description: "Aset elektronik telah berhasil didaftarkan.",
    });

    router.push('/electronic');
    form.resetForm();
    currentStep.value = 1;
  } catch (error) {
    console.error("Error submitting form:", error);


    toast({
      title: "Gagal Mendaftarkan Aset",
      description: "Terjadi kesalahan saat mendaftarkan aset elektronik.",
      variant: "destructive"
    });
  }
};

const onNextStep = () => {
  currentStep.value += 1;
};

const onPrevStep = () => {
  currentStep.value -= 1;
};

const submitButtonText = computed(() => {
  return currentStep.value === 2 ? "Daftarkan Aset" : "Selanjutnya";
});

const submitButtonClasses = computed(() => {
  return currentStep.value === 2
    ? "px-6 py-2 bg-green-600 text-white hover:bg-green-700"
    : "px-6 py-2 bg-blue-600 text-white hover:bg-blue-700";
});
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <div class="bg-white shadow-lg rounded-lg p-8">
      <h1 class="text-xl font-bold text-gray-800 mb-8 text-start">
        Registrasi Aset Elektronik
      </h1>

      <!-- Progress Indicator -->
      <div class="flex justify-center mb-6">
        <div class="flex items-center">
          <div class="w-10 h-10 rounded-full flex items-center justify-center mr-2"
            :class="currentStep === 1 ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-600'">
            1
          </div>
          <div class="w-24 h-1 bg-gray-200 mx-2"></div>
          <div class="w-10 h-10 rounded-full flex items-center justify-center"
            :class="currentStep === 2 ? 'bg-blue-600 text-white' : 'bg-gray-200 text-gray-600'">
            2
          </div>
        </div>
      </div>

      <Form @submit="onSubmit" :form="form">
        <!-- Step 1: Section Fields -->
        <div v-if="currentStep == 1" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <FormField v-for="field in step1Fields" :key="field" :name="field" v-slot="{ field: formField, error }">
            <FormItem class="mb-4">
              <FormLabel class="block text-sm font-medium text-gray-700 mb-2">
                {{ getFieldLabel(field) }}
              </FormLabel>
              <FormControl>

                <Select v-if="[
                  'namasupllier',
                  'tipeasset',
                  'kondisiasset',
                  'status'
                ].includes(field)" v-bind="formField">
                  <SelectTrigger class="w-full">
                    <SelectValue :placeholder="`Pilih ${getFieldLabel(field)}`" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem v-for="option in getOptionsForField(field)" :key="option.value" :value="option.value">
                      {{ option.label }}
                    </SelectItem>
                  </SelectContent>
                </Select>

                <Input v-else-if="['tandaterima', 'hasilpemeriksaan'].includes(field)" type="file" class="w-full"
                  v-bind="formField" />

                <Input v-else-if="['achievedDate'].includes(field)" type="date" class="w-full" v-bind="formField" />

                <Input v-else-if="['nilaiasset'].includes(field)" type="number" class="w-full" v-bind="formField" />

                <Input v-else type="text" class="w-full" v-bind="formField" />
              </FormControl>
              <FormMessage class="text-red-500 text-xs mt-1">{{ error }}</FormMessage>
            </FormItem>
          </FormField>
        </div>


        <!-- Step 2: Section Fields -->
        <div v-if="currentStep == 2" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <FormField v-for="field in step2Fields" :key="field" :name="field" v-slot="{ field: formField, error }">
            <FormItem class="mb-4">
              <FormLabel class="block text
              -sm font-medium text-gray-700 mb-2">
                {{ getFieldLabel(field) }}
              </FormLabel>
              <FormControl>

                <Select v-if="[
                  'penanggungjawabasset',
                  'divpengguna',
                  'jangkamasapakaiaset'
                ].includes(field)" v-bind="formField">
                  <SelectTrigger class="w-full">
                    <SelectValue :placeholder="`Pilih ${getFieldLabel(field)}`" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem v-for="option in getOptionsForField(field)" :key="option.value" :value="option.value">
                      {{ option.label }}
                    </SelectItem>
                  </SelectContent>
                </Select>

                <Input v-else-if="['notapembelian'].includes(field)" type="file" class="w-full" v-bind="formField" />

                <Input
                  v-else-if="['tanggalaktivasi', 'tanggalasetkeluar', 'garansimulai', 'garansiakhir'].includes(field)"
                  type="date" class="w-full" v-bind="formField" />

                <Input v-else-if="['ram', 'rom'].includes(field)" type="number" class="w-full" v-bind="formField" />

                <Input v-else type="text" class="w-full" v-bind="formField" />
              </FormControl>
              <FormMessage class="text-red-500 text-xs mt-1">{{ error }}</FormMessage>
            </FormItem>
          </FormField>
        </div>
        <div class="flex justify-between mt-8 space-x-4">
          <Button v-if="currentStep > 1" variant="outline" type="button" class="px-6 py-2" @click="onPrevStep">
            Kembali
          </Button>

          <div class="flex-grow"></div>
          <Button v-if="currentStep < 2" variant="outline" type="button"
            class="px-6 py-2 bg-blue-600 text-white hover:bg-blue-700" @click="onNextStep">
            Selanjutnya
          </Button>

          <Button v-if="currentStep === 2" type="submit" :class="submitButtonClasses">
            {{ submitButtonText }}
          </Button>
        </div>
      </Form>
    </div>
  </div>
</template>
