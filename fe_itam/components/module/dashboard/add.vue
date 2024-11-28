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
import { Input } from "@/components/ui/input";

const props = defineProps<{
  type: string;
  data: any[];
}>();

/* handle form */
const formSchema = toTypedSchema(
  z.object({
    totalRevenue: z.number().min(0),
    subscriptions: z.number().min(0),
    sales: z.number().min(0),
    activeNow: z.number().min(0),
    dateRange: z.string().refine((v) => v, { message: "Date is required." }),
    overview: z.string().min(1),
  })
);

const { handleSubmit, setFieldValue, values } = useForm({
  validationSchema: formSchema,
  initialValues: {
    totalRevenue: 0,
    subscriptions: 0,
    sales: 0,
    activeNow: 0,
    dateRange: "",
    overview: "",
  },
});

const onSubmit = handleSubmit(async (values) => {
  try {
    const formData = new FormData();
    formData.append("totalRevenue", values.totalRevenue.toString());
    formData.append("subscriptions", values.subscriptions.toString());
    formData.append("sales", values.sales.toString());
    formData.append("activeNow", values.activeNow.toString());
    formData.append("dateRange", values.dateRange);
    formData.append("overview", values.overview);
    q;
    const response = await $fetch(
      `${process.env.API_URL}/total/asset-aplikasi`,
      {
        method: "GET",
        body: formData,
      }
    );

    console.log("Data berhasil disimpan:", response);
    navigateTo("/dashboard");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
});
</script>

<template>
  <div class="p-8 bg-white shadow-lg rounded-lg">
    <h1 class="text-2xl font-bold mb-6">Dashboard</h1>

    <form @submit="onSubmit">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="totalRevenue">
          <FormItem>
            <FormLabel>Total Revenue</FormLabel>
            <FormControl>
              <Input type="number" v-model="values.totalRevenue" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="subscriptions">
          <FormItem>
            <FormLabel>Subscriptions</FormLabel>
            <FormControl>
              <Input type="number" v-model="values.subscriptions" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="sales">
          <FormItem>
            <FormLabel>Sales</FormLabel>
            <FormControl>
              <Input type="number" v-model="values.sales" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="activeNow">
          <FormItem>
            <FormLabel>Active Now</FormLabel>
            <FormControl>
              <Input type="number" v-model="values.activeNow" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="dateRange">
          <FormItem>
            <FormLabel>Date Range</FormLabel>
            <FormControl>
              <Input type="date" v-model="values.dateRange" />
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>

        <FormField name="overview">
          <FormItem>
            <FormLabel>Overview</FormLabel>
            <FormControl>
              <Select v-model="values.overview">
                <SelectTrigger>
                  <SelectValue placeholder="Select Overview" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="daily">Daily</SelectItem>
                  <SelectItem value="weekly">Weekly</SelectItem>
                  <SelectItem value="monthly">Monthly</SelectItem>
                </SelectContent>
              </Select>
            </FormControl>
            <FormMessage />
          </FormItem>
        </FormField>
      </div>

      <Button type="submit" class="mt-4">Submit</Button>
    </form>
  </div>
</template>
