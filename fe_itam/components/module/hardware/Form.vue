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
  data?: any;
}>();

const config = useRuntimeConfig();
const { showLoading, hideLoading } = useLoading();

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    supplier: z.string().min(1, "Supplier harus dipilih"),
    serial_number: z.string().min(1, "Serial number harus diisi"),
    asset_type: z.string().min(1, "Tipe aset harus dipilih"),
    nota_number: z.string().min(1, "Nomor nota harus diisi"),
    achieved_date: z.string().min(1, "Tanggal penerimaan harus diisi"),
    out_date: z.string().min(1, "Tanggal keluar harus diisi"),
    brand: z.string().min(1, "Merek harus diisi"),
    model: z.string().min(1, "Model harus diisi"),
    warranty_number: z.string().min(1, "Nomor garansi harus diisi"),
    warranty_start: z.string().min(1, "Tanggal mulai garansi harus diisi"),
    warranty_end: z.string().min(1, "Tanggal akhir garansi harus diisi"),
    activation_date: z.string().min(1, "Tanggal aktivasi harus diisi"),
    usage_period: z.string().min(1, "Jangka masa pakai harus dipilih"),
    pic: z.string().min(1, "Penanggung jawab harus dipilih"),
    asset_status: z.string().min(1, "Status aset harus dipilih"),
    inspection_result: z.string().min(1, "Hasil pemeriksaan harus diisi"),
    storage_location: z.string().min(1, "Lokasi penyimpanan harus diisi"),
    purchase_note: z.any(),
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
      body: formData
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

/* hold datefield values */
const achievedDate = computed({
  get: () =>
    values.achieved_date ? parseDate(values.achieved_date) : undefined,
  set: (val) => val,
});

const outDate = computed({
  get: () => (values.out_date ? parseDate(values.out_date) : undefined),
  set: (val) => val,
});

const warrantyStart = computed({
  get: () =>
    values.warranty_start ? parseDate(values.warranty_start) : undefined,
  set: (val) => val,
});

const warrantyEnd = computed({
  get: () => (values.warranty_end ? parseDate(values.warranty_end) : undefined),
  set: (val) => val,
});

const activationDate = computed({
  get: () =>
    values.activation_date ? parseDate(values.activation_date) : undefined,
  set: (val) => val,
});
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Registrasi Aset Hardware</h1>
    <form class="grid grid-cols-2 gap-x-4 gap-y-3">
      <FormField v-slot="{ componentField }" name="supplier">
        <FormItem>
          <FormLabel>Nama Supplier</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Supplier" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in suppliers">
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

      <FormField v-slot="{ componentField }" name="serial_number">
        <FormItem>
          <FormLabel>Serial Number</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="asset_type">
        <FormItem>
          <FormLabel>Tipe Asset</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Tipe" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in assetTypes">
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

      <FormField v-slot="{ componentField }" name="nota_number">
        <FormItem>
          <FormLabel>Nomor Nota</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="achieved_date">
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
                      !achievedDate && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    achievedDate
                      ? df.format(toDate(achievedDate))
                      : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="achievedDate"
                @update:model-value="
                  (v) => setFieldValue('achieved_date', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="out_date">
        <FormItem class="flex flex-col">
          <FormLabel>Tanggal Keluar</FormLabel>
          <Popover>
            <PopoverTrigger as-child>
              <FormControl>
                <Button
                  variant="outline"
                  :class="
                    cn(
                      'ps-3 text-start font-normal',
                      !outDate && 'text-muted-foreground'
                    )
                  "
                >
                  <span>{{
                    outDate ? df.format(toDate(outDate)) : "Pilih tanggal"
                  }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="outDate"
                @update:model-value="
                  (v) => setFieldValue('out_date', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="brand">
        <FormItem>
          <FormLabel>Merek</FormLabel>
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

      <FormField v-slot="{ componentField }" name="warranty_number">
        <FormItem>
          <FormLabel>Nomor Kartu Garansi</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="warranty_start">
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
                  (v) => setFieldValue('warranty_start', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="warranty_end">
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
                  (v) => setFieldValue('warranty_end', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField name="activation_date">
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
                  (v) => setFieldValue('activation_date', v?.toString())
                "
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="usage_period">
        <FormItem>
          <FormLabel>Jangka Masa Pakai</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Jangka Waktu" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in usagePeriods">
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

      <FormField v-slot="{ componentField }" name="pic">
        <FormItem>
          <FormLabel>Penanggung Jawab</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Penanggung Jawab" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in pics">
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

      <FormField v-slot="{ componentField }" name="asset_status">
        <FormItem>
          <FormLabel>Status Asset</FormLabel>
          <Select v-bind="componentField">
            <FormControl>
              <SelectTrigger>
                <SelectValue placeholder="Pilih Status" />
              </SelectTrigger>
            </FormControl>
            <SelectContent>
              <SelectGroup>
                <template v-for="item in assetStatuses">
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

      <FormField v-slot="{ componentField }" name="inspection_result">
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

      <FormField v-slot="{ componentField }" name="storage_location">
        <FormItem>
          <FormLabel>Lokasi Penyimpanan</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="purchase_note">
        <FormItem>
          <FormLabel>Nota Pembelian</FormLabel>
          <FormControl>
            <Input
              type="file"
              accept=".pdf,.jpg,.jpeg,.png"
              @change="(e) => setFieldValue('purchase_note', e.target.files[0])"
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
