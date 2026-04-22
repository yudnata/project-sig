<script setup lang="ts">
import { useMapPointsStore } from '@/stores/mapPoints'
import MapContainer from '@/components/Map/MapContainer.vue'

const store = useMapPointsStore()
store.isEditMode = false
store.filterMyPoints = false
</script>

<template>
  <div class="h-full w-full relative">
    <MapContainer />
  </div>
</template>
