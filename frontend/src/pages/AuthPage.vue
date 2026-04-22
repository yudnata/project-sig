<script setup lang="ts">
import { useRouter } from 'vue-router'
import { googleTokenLogin } from 'vue3-google-login'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notifications'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const loginWithGoogle = async () => {
  try {
    // Opens a popup — no One Tap, no origin propagation delay
    const tokenResponse = await googleTokenLogin()

    // Get user info using the access token
    const userInfoRes = await fetch('https://www.googleapis.com/oauth2/v3/userinfo', {
      headers: { Authorization: `Bearer ${tokenResponse.access_token}` }
    })
    const userInfo = await userInfoRes.json()

    // Send to backend
    const res = await fetch(`${API_URL}/auth/sso`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: userInfo.email,
        name: userInfo.name,
        sso_provider: 'google',
        sso_id: userInfo.sub,
        avatar_url: userInfo.picture,
      })
    })

    const result = await res.json()
    if (result.success) {
      authStore.login(result.data.token, result.data.user)
      notificationStore.success('Selamat datang kembali!')
      router.push('/dashboard')
    } else {
      notificationStore.error(result.message || 'Login SSO gagal di sisi server.')
    }
  } catch (err) {
    console.error('SSO Error:', err)
    notificationStore.error('Terjadi kesalahan saat login Google.')
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50 -mt-20">
    <div class="bg-white p-8 rounded-xl shadow-[0_4px_24px_rgba(0,0,0,0.06)] w-full max-w-sm">
      <h1 class="text-2xl font-bold text-center text-primary mb-2">Masuk</h1>
      <p class="text-sm text-center text-gray-500 mb-6">Akses kontrol SIG Tata Ruang Anda</p>

      <div class="space-y-4">

        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Username</label>
          <input type="text" class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="Masukkan username">
        </div>
        <div>
          <label class="block text-xs font-semibold text-gray-600 uppercase mb-1">Password</label>
          <input type="password" class="w-full px-4 py-2 bg-gray-100 border-transparent rounded-lg focus:bg-white focus:border-primary focus:ring-2 focus:ring-primary/20 transition-all text-sm"
            placeholder="••••••••">
        </div>
        <button class="w-full bg-primary text-white font-bold py-2.5 rounded-lg hover:bg-primary-hover active:scale-[0.98] transition-all">
          Masuk Sekarang
        </button>

        <div class="relative flex py-2 items-center">
          <div class="flex-grow border-t border-gray-200"></div>
          <span class="flex-shrink-0 mx-4 text-gray-400 text-xs">ATAU</span>
          <div class="flex-grow border-t border-gray-200"></div>
        </div>

        <!-- Custom Google Button -->
        <button @click="loginWithGoogle"
          class="w-full flex items-center justify-center gap-3 px-4 py-2.5 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 active:scale-[0.98] transition-all shadow-sm">
          <!-- Google G Logo SVG -->
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" class="w-5 h-5 shrink-0">
            <path fill="#FFC107"
              d="M43.6 20.1H42V20H24v8h11.3C33.7 32.7 29.2 36 24 36c-6.6 0-12-5.4-12-12s5.4-12 12-12c3.1 0 5.8 1.2 7.9 3l5.7-5.7C34 6.5 29.3 4 24 4 12.9 4 4 12.9 4 24s8.9 20 20 20 20-8.9 20-20c0-1.3-.1-2.7-.4-3.9z" />
            <path fill="#FF3D00" d="m6.3 14.7 6.6 4.8C14.7 16 19 13 24 13c3.1 0 5.8 1.2 7.9 3l5.7-5.7C34 6.5 29.3 4 24 4 16.3 4 9.7 8.3 6.3 14.7z" />
            <path fill="#4CAF50" d="M24 44c5.2 0 9.9-2 13.4-5.2l-6.2-5.2C29.3 35.5 26.8 36 24 36c-5.2 0-9.6-3.3-11.3-7.9l-6.5 5C9.5 39.6 16.2 44 24 44z" />
            <path fill="#1976D2" d="M43.6 20.1H42V20H24v8h11.3c-.8 2.3-2.3 4.2-4.3 5.6l6.2 5.2C36.9 37.3 44 32 44 24c0-1.3-.1-2.7-.4-3.9z" />
          </svg>
          <span class="text-sm font-semibold text-gray-700">Masuk dengan Google</span>
        </button>
      </div>
    </div>
  </div>
</template>
