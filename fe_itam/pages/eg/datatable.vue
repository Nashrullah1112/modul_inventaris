<script setup lang="ts">
import { columns } from "@/components/columns";
import type { Payment } from "@/components/columns";
import DataTable from "@/components/DataTable.vue";

const data = ref<Payment[]>([]);

async function getData(): Promise<Payment[]> {
  // Fetch data from your API here.
  return [
    {
      id: "a1b2c3d4",
      amount: 150,
      status: "approved",
      email: "user1@example.com",
    },
    {
      id: "e5f6g7h8",
      amount: 200,
      status: "pending",
      email: "user2@example.com",
    },
    {
      id: "i9j0k1l2",
      amount: 250,
      status: "rejected",
      email: "user3@example.com",
    },
    {
      id: "m3n4o5p6",
      amount: 300,
      status: "approved",
      email: "user4@example.com",
    },
    {
      id: "q7r8s9t0",
      amount: 350,
      status: "pending",
      email: "user5@example.com",
    },
    {
      id: "u1v2w3x4",
      amount: 400,
      status: "rejected",
      email: "user6@example.com",
    },
    {
      id: "y5z6a7b8",
      amount: 450,
      status: "approved",
      email: "user7@example.com",
    },
    {
      id: "c9d0e1f2",
      amount: 500,
      status: "pending",
      email: "user8@example.com",
    },
    {
      id: "g3h4i5j6",
      amount: 550,
      status: "rejected",
      email: "user9@example.com",
    },
    {
      id: "k7l8m9n0",
      amount: 600,
      status: "approved",
      email: "user10@example.com",
    },
    {
      id: "o1p2q3r4",
      amount: 650,
      status: "pending",
      email: "user11@example.com",
    },
    {
      id: "s5t6u7v8",
      amount: 700,
      status: "rejected",
      email: "user12@example.com",
    },
    {
      id: "w9x0y1z2",
      amount: 750,
      status: "approved",
      email: "user13@example.com",
    },
    {
      id: "a3b4c5d6",
      amount: 800,
      status: "pending",
      email: "user14@example.com",
    },
    {
      id: "e7f8g9h0",
      amount: 850,
      status: "rejected",
      email: "user15@example.com",
    },
  ];
}

onMounted(async () => {
  data.value = await getData();
});
</script>

<template>
  <div>
    <DataTable :columns="columns" :data="data" />
  </div>
</template>
