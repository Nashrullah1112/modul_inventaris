<script setup lang="ts">
import { useToast } from '@/components/ui/toast/use-toast'

definePageMeta({
  layout: 'plain'
})

const config = useRuntimeConfig()
const router = useRouter()
const { toast } = useToast()

const username = ref("");
const password = ref("");

const handleLogin = async () => {
  try {
    const { data, status, error } = await useFetch(config.public.API_URL + '/login', {
      method: 'POST',
      body: {
        username: username.value,
        password: password.value,
      }
    })

    if (status.value == 'success' && data.value?.data?.hasOwnProperty('Token')) {
      if (process.client) {
        localStorage.setItem('token', data.value.data.Token);
        router.push('/');
      }
    } else {
      toast({
        title: 'Failed',
        description: `Failed to log in. ${data.value?.message || ''}`,
      })
    }
  } catch (err) {
    console.error("Error occured:", err);
  }
}
</script>

<template>
  <div class="min-h-screen flex">
    <!-- Bagian Kiri - Gambar -->
    <div class="hidden lg:flex lg:w-1/2 bg-[#1a237e]">
      <div class="relative w-full h-full">
        <img
          src="https://images.unsplash.com/photo-1484417894907-623942c8ee29?q=80&w=1932&auto=format&fit=crop"
          alt="Office"
          class="w-full h-full object-cover opacity-70"
        />
        <div
          class="absolute inset-0 flex flex-col justify-center items-center text-white p-12"
        >
          <h1 class="text-4xl font-bold mb-4">IT Asset Management</h1>
          <p class="text-xl text-center">
            Sistem manajemen aset IT yang terintegrasi untuk monitoring dan
            kontrol yang lebih efektif
          </p>
        </div>
      </div>
    </div>

    <!-- Bagian Kanan - Form Login -->
    <div
      class="w-full lg:w-1/2 flex items-center justify-center p-8 bg-gray-50"
    >
      <div class="w-full max-w-md">
        <div class="text-center mb-8">
          <h2 class="text-3xl font-bold text-[#1a237e]">Selamat Datang</h2>
          <p class="text-gray-600 mt-2">Silakan masuk ke akun Anda</p>
        </div>

        <form @submit.prevent="handleLogin" class="space-y-6">
          <div class="space-y-2">
            <label
              for="username"
              class="text-sm font-medium leading-none text-gray-700"
            >
              Username
            </label>
            <input
              id="username"
              v-model="username"
              type="text"
              class="flex h-10 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-2 focus:ring-[#1a237e] focus:border-transparent"
              placeholder="Masukkan username"
              required
            />
          </div>

          <div class="space-y-2">
            <label
              for="password"
              class="text-sm font-medium leading-none text-gray-700"
            >
              Password
            </label>
            <input
              id="password"
              v-model="password"
              type="password"
              class="flex h-10 w-full rounded-md border border-gray-300 bg-white px-3 py-2 text-sm shadow-sm transition-colors focus:outline-none focus:ring-2 focus:ring-[#1a237e] focus:border-transparent"
              placeholder="Masukkan password"
              required
            />
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember"
                type="checkbox"
                class="h-4 w-4 rounded border-gray-300 text-[#1a237e] focus:ring-[#1a237e]"
              />
              <label for="remember" class="ml-2 text-sm text-gray-600">
                Ingat saya
              </label>
            </div>
            <a href="#" class="text-sm text-[#1a237e] hover:underline">
              Lupa password?
            </a>
          </div>

          <button
            type="submit"
            class="w-full bg-[#1a237e] text-white h-11 px-4 py-2 rounded-md hover:bg-[#283593] transition-colors focus:outline-none focus:ring-2 focus:ring-[#1a237e] focus:ring-offset-2 font-medium"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  </div>
</template>
