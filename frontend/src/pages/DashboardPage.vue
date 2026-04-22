<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import { useNotificationStore } from '@/stores/notifications'
import MapContainer from '@/components/Map/MapContainer.vue'

const store = useMapPointsStore()
const notificationStore = useNotificationStore()

const handleMapClick = (data: { lat: number; lng: number; address?: string }) => {
  notificationStore.info('Membuka form...')
  store.openModal({
    id: 0,
    name: '',
    latitude: data.lat,
    longitude: data.lng,
    type_id: 1,
    address: data.address || '',
  })
}
</script>

<template>
  <div class="h-full w-full flex flex-col overflow-hidden relative">
    <MapContainer @map-clicked="handleMapClick" />
  </div>
</template>
