<template>
  <div
    class="p-8 bg-white shadow-lg rounded-lg"
    :class="isOpen ? 'lg:ml-64' : ''"
  >
    <h1 class="text-2xl font-bold mb-6">Dashboard</h1>

    <Form @submit="registerDashboard">
      <div class="grid grid-cols-2 gap-x-4">
        <FormField name="totalRevenue">
          <FormItem>
            <FormLabel>Total Revenue</FormLabel>
            <FormControl>
              <Input type="number" v-model="totalRevenue" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="subscriptions">
          <FormItem>
            <FormLabel>Subscriptions</FormLabel>
            <FormControl>
              <Input type="number" v-model="subscriptions" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="sales">
          <FormItem>
            <FormLabel>Sales</FormLabel>
            <FormControl>
              <Input type="number" v-model="sales" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="activeNow">
          <FormItem>
            <FormLabel>Active Now</FormLabel>
            <FormControl>
              <Input type="number" v-model="activeNow" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="dateRange">
          <FormItem>
            <FormLabel>Date Range</FormLabel>
            <FormControl>
              <Input type="date" v-model="dateRange" />
            </FormControl>
          </FormItem>
        </FormField>

        <FormField name="overview">
          <FormItem>
            <FormLabel>Overview</FormLabel>
            <FormControl>
              <Select v-model="overview">
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
          </FormItem>
        </FormField>
      </div>
    </Form>
  </div>
</template>

<script setup>
import { useState } from "#app";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const baseURL = "http://localhost:3000/api";

const isOpen = useState("is-sidebar-open");
const router = useRouter();

const totalRevenue = ref(0);
const subscriptions = ref(0);
const sales = ref(0);
const activeNow = ref(0);
const dateRange = ref("");
const overview = ref("");

const registerDashboard = async () => {
  try {
    const formData = new FormData();
    formData.append("totalRevenue", totalRevenue.value);
    formData.append("subscriptions", subscriptions.value);
    formData.append("sales", sales.value);
    formData.append("activeNow", activeNow.value);
    formData.append("dateRange", dateRange.value);
    formData.append("overview", overview.value);

    const response = await $fetch(`${baseURL}/dashboard`, {
      method: "POST",
      body: formData,
    });

    console.log("Data berhasil disimpan:", response);
    router.push("/dashboard");
  } catch (error) {
    console.error("Terjadi kesalahan:", error);
  }
};
</script>
