<script setup lang="ts">
import { cn } from '@/lib/utils'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'
import { toDate } from 'radix-vue/date'
import { CalendarDate, DateFormatter, getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { CalendarIcon } from '@radix-icons/vue'

import { Button } from '@/components/ui/button'
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Calendar } from '@/components/ui/calendar'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Textarea } from '@/components/ui/textarea'

const props = defineProps<{
  type: string,
  data?: any,
}>()

const config = useRuntimeConfig()
const { showLoading, hideLoading } = useLoading()

/* handle form */
const formSchema = toTypedSchema(z.object({
  supplier: z.string().min(1),
  asset_type: z.string().min(1),
  nota_number: z.string().min(1),
  achieved_date: z
    .string()
    .refine(v => v, { message: 'Date is required.' }),
  serial_number: z.string().min(1),
  brand: z.string().min(1),
  model: z.string().min(1),
  specification: z.string().min(1),
}))

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
})

const onSubmit = handleSubmit(async (values) => {
  showLoading()

  const { data, status } = await useFetch(config.public.API_URL_2 + '/t_asset_hardware', {
    method: 'POST',
    body: values,
    headers: { 
      apiKey: config.public.API_KEY_2,
    },
  })

  hideLoading()

  if (status.value == 'success') {
    navigateTo('/hardware')
  }
})

const df = new DateFormatter('en-US', {
  dateStyle: 'long',
})

/* hold datefield value */
const achievedDate = computed({
  get: () => values.achieved_date ? parseDate(values.achieved_date) : undefined,
  set: val => val,
})

/* Dummy data */
const suppliers = [
  { label: 'PT Surya Kencana', value: '1' },
  { label: 'PT DJAJA PUTRA', value: '2' },
  { label: 'PT INDO IT', value: '3' },
]
const assetTypes = [
  { label: 'Processor', value: '1' },
  { label: 'Motherboard', value: '2' },
  { label: 'RAM', value: '3' },
  { label: 'Storage', value: '4' },
  { label: 'Power Supply', value: '5' },
  { label: 'Casing', value: '6' },
  { label: 'Cooling Fan', value: '7' },
]
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

      <FormField v-slot="{ componentField }" name="nota_number" >
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
                  variant="outline" :class="cn(
                    'ps-3 text-start font-normal',
                    !achievedDate && 'text-muted-foreground',
                  )"
                >
                  <span>{{ achievedDate ? df.format(toDate(achievedDate)) : "Pick a date" }}</span>
                  <CalendarIcon class="ms-auto h-4 w-4 opacity-50" />
                </Button>
                <input hidden>
              </FormControl>
            </PopoverTrigger>
            <PopoverContent class="w-auto p-0">
              <Calendar
                v-model="achievedDate"
                calendar-label="Date of birth"
                @update:model-value="(v) => {
                  if (v) {
                    setFieldValue('achieved_date', v.toString())
                  }
                  else {
                    setFieldValue('achieved_date', undefined)
                  }
                }"
              />
            </PopoverContent>
          </Popover>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="serial_number" >
        <FormItem>
          <FormLabel>Serial Number</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="brand" >
        <FormItem>
          <FormLabel>Merek</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="model" >
        <FormItem>
          <FormLabel>Model</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>

      <FormField v-slot="{ componentField }" name="specification" >
        <FormItem>
          <FormLabel>Spesifikasi</FormLabel>
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
    </form>

    <div class="flex justify-end mt-4 space-x-2">
      <Button @click="onSubmit">
        Submit
      </Button>
    </div>
  </div>
</template>
