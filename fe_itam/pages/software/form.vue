<script setup lang="ts">
import { useForm } from "vee-validate";
import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

import { Button } from "@/components/ui/button";
import {
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";

const formSchema = toTypedSchema(
  z.object({
    vendorName: z.string().min(2).max(50),
    applicationName: z.string().min(2).max(50),
    applicationType: z.string().min(2).max(50),
    applicationUrl: z.string().min(2).max(50),
    certificationType: z.string().min(2).max(50),
    creationDate: z.string().min(2).max(50),
    acceptedDate: z.string().min(2).max(50),
    serverLocation: z.string().min(2).max(50),
    domainActiveDate: z.string().min(2).max(50),
    domainExpiredDate: z.string().min(2).max(50),
  })
);

const form = useForm({
  validationSchema: formSchema,
});

const onSubmit = form.handleSubmit((values) => {
  console.log("Form submitted!", values);
});
</script>

<template>
  <form @submit="onSubmit">
    <div class="grid lg:grid-cols-2 gap-x-4">
      <FormField v-slot="{ componentField }" name="vendorName">
        <FormItem>
          <FormLabel>Nama Vendor</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="applicationName">
        <FormItem>
          <FormLabel>Nama Aplikasi</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="applicationType">
        <FormItem>
          <FormLabel>Tipe Aplikasi</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="applicationUrl">
        <FormItem>
          <FormLabel>Url Aplikasi</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="certificationType">
        <FormItem>
          <FormLabel>Tipe Sertifikasi Asset </FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="creationDate">
        <FormItem>
          <FormLabel>Tanggal Pembuatan </FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="acceptedDate">
        <FormItem>
          <FormLabel>Tanggal Penerimaan</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="serverLocation">
        <FormItem>
          <FormLabel>Lokasi Server Penyimpanan</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="domainActiveDate">
        <FormItem>
          <FormLabel>Tanggal Aktif Domain</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
      <FormField v-slot="{ componentField }" name="domainExpiredDate">
        <FormItem>
          <FormLabel>Tanggal Expired Domain</FormLabel>
          <FormControl>
            <Input type="text" placeholder="" v-bind="componentField" />
          </FormControl>
          <FormMessage />
        </FormItem>
      </FormField>
    </div>

    <Button type="submit"> Submit </Button>
  </form>
</template>
