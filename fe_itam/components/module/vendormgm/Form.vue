<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import {
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const props = defineProps<{
  type: string;
  data?: any;
}>();

const config = useRuntimeConfig()
const route = useRoute()
const { showLoading, hideLoading } = useLoading()
const { toast } = useToast()

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    nama_pic: z.string().min(1),
    email: z.string().email(),
    nomor_kontak: z.string().min(10),
    lokasi_perusahaan: z.string().min(1),
    nomor_siup: z.string().optional(),
    nomor_nib: z.string().optional(),
    nomor_npwp: z.string().optional(),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
});

const dataId = route.params.id
const endpoint = `/vendor${props.type == 'new' ? '' : '/' + dataId}`

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
      navigateTo('/vendormgm')
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
      
      setFieldValue('nama_pic', exData.nama_pic)
      setFieldValue('email', exData.email)
      setFieldValue('nomor_kontak', exData.nomor_kontak)
      setFieldValue('lokasi_perusahaan', exData.lokasi_perusahaan)
      setFieldValue('nomor_siup', exData.nomor_siup)
      setFieldValue('nomor_nib', exData.nomor_nib)
      setFieldValue('nomor_npwp', exData.nomor_npwp)
    }
  } catch (error) {
    console.error("Error occured:", error);
  }

  hideLoading()
}

if (props.type == 'edit') {
  getExistingData()
}
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">
      {{ props.type == 'new' ? 'Tambah' : 'Edit' }} Vendor
    </h1>

    <form>
      <div class="grid grid-cols-2 gap-x-4 gap-y-2">
        <FormField v-slot="{ componentField }" name="nama_pic">
          <FormItem>
            <FormLabel>Nama PIC</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="email">
          <FormItem>
            <FormLabel>Email</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_kontak">
          <FormItem>
            <FormLabel>Nomor Kontak</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="lokasi_perusahaan">
          <FormItem>
            <FormLabel>Lokasi Perusahaan</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_siup">
          <FormItem>
            <FormLabel>Nomor SIUP</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_nib">
          <FormItem>
            <FormLabel>Nomor NIB</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField v-slot="{ componentField }" name="nomor_npwp">
          <FormItem>
            <FormLabel>Nomor NPWP</FormLabel>
            <FormControl>
              <Input type="text" v-bind="componentField" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>
    </form>
    
    <div class="flex justify-end mt-4 space-x-2">
      <Button variant="outline" @click="navigateTo('/vendormgm')">Cancel</Button>
      <Button @click="onSubmit">Submit</Button>
    </div>
  </div>
</template>
