<script setup lang="ts">
import { RouterView, RouterLink, useRoute } from 'vue-router'
import { useFacilityStore } from '@/stores/facility'
import { onMounted, computed } from 'vue'

const store = useFacilityStore()
onMounted(() => {
  store.initFacilities()
})

const route = useRoute()

const pageTitle = computed(() => {
  if (route.path === '/') return 'Dashboard'
  if (route.path === '/add') return 'Tambah Fasilitas Baru'
  if (route.path === '/list') return 'Daftar Fasilitas Kesehatan'
  if (route.path.startsWith('/edit')) return 'Edit Fasilitas'
  return 'Sistem Informasi Geografis'
})
</script>

<template>
  <div class="flex h-screen bg-surface font-sans text-text antialiased overflow-hidden">
    <!-- Sidebar Desktop -->
    <aside class="w-64 bg-primary text-white flex-col hidden md:flex shrink-0">
      <div class="p-6">
        <h1 class="text-xl font-bold leading-tight">SIG Faskes</h1>
        <p class="text-surface/80 text-xs mt-1">Sistem Pemetaan Bali</p>
      </div>
      <nav class="flex-1 px-4 space-y-2 mt-4">
        <RouterLink
          to="/"
          :class="[
            `flex items-center gap-3 px-4 py-2.5 rounded transition`,
            route.path === `/`
              ? `bg-white/20 text-white font-bold`
              : `text-white/80 hover:bg-white/10 hover:text-white`,
          ]"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M2.25 12l8.954-8.955c.44-.439 1.152-.439 1.591 0L21.75 12M4.5 9.75v10.125c0 .621.504 1.125 1.125 1.125H9.75v-4.875c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125V21h4.125c.621 0 1.125-.504 1.125-1.125V9.75M8.25 21h8.25"
            />
          </svg>
          <span class="font-medium">Dashboard</span>
        </RouterLink>

        <RouterLink
          to="/add"
          :class="[
            `flex items-center gap-3 px-4 py-2.5 rounded transition`,
            route.path === `/add` || route.path.startsWith(`/edit`)
              ? `bg-white/20 text-white font-bold`
              : `text-white/80 hover:bg-white/10 hover:text-white`,
          ]"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 9v6m3-3H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
            />
          </svg>
          <span class="font-medium">Tambah Fasilitas</span>
        </RouterLink>

        <RouterLink
          to="/list"
          :class="[
            `flex items-center gap-3 px-4 py-2.5 rounded transition`,
            route.path === `/list`
              ? `bg-white/20 text-white font-bold`
              : `text-white/80 hover:bg-white/10 hover:text-white`,
          ]"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="w-5 h-5"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M8.25 6.75h12M8.25 12h12m-12 5.25h12M3.75 6.75h.007v.008H3.75V6.75zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zM3.75 12h.007v.008H3.75V12zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0zm-.375 5.25h.007v.008H3.75v-.008zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z"
            />
          </svg>
          <span class="font-medium">Daftar Fasilitas</span>
        </RouterLink>
      </nav>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 flex flex-col h-screen overflow-hidden relative">
      <!-- Mobile Header & Nav -->
      <header class="bg-primary text-white shadow-sm md:hidden p-4 flex flex-col gap-2">
        <div>
          <h1 class="text-lg font-bold">SIG Faskes Bali</h1>
        </div>
        <nav class="flex overflow-auto text-sm space-x-2 pb-1">
          <RouterLink
            to="/"
            class="px-3 py-1.5 rounded text-white/80"
            active-class="bg-white/20 !text-white font-bold"
            >Dashboard</RouterLink
          >
          <RouterLink
            to="/add"
            class="px-3 py-1.5 rounded shrink-0 text-white/80"
            active-class="bg-white/20 !text-white font-bold"
            >Tambah Data</RouterLink
          >
          <RouterLink
            to="/list"
            class="px-3 py-1.5 rounded shrink-0 text-white/80"
            active-class="bg-white/20 !text-white font-bold"
            >Daftar Data</RouterLink
          >
        </nav>
      </header>

      <!-- Desktop Header -->
      <header
        class="hidden md:flex bg-white border-b border-text/10 h-16 items-center px-6 shrink-0 shadow-sm z-10"
      >
        <h2 class="text-lg font-bold text-text">{{ pageTitle }}</h2>
      </header>

      <!-- Scrollable Main Area -->
      <main class="flex-1 overflow-auto p-4 md:p-6 relative bg-surface bg-grid-pattern">
        <RouterView />
      </main>
    </div>
  </div>
</template>
