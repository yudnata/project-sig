<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, nextTick, h, render } from 'vue'
import L from 'leaflet'
import 'leaflet/dist/leaflet.css'
import { useMapPointsStore } from '@/stores/mapPoints'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
import MapLegend from './MapLegend.vue'
import MapPopupContent from './MapPopupContent.vue'

const emit = defineEmits<{
  (e: 'map-clicked', data: { lat: number; lng: number; address?: string }): void
}>()

const store = useMapPointsStore()
const authStore = useAuthStore()
const { filteredPoints, objectTypes, isEditMode, isSidebarExpanded } = storeToRefs(store)

const mapContainer = ref<HTMLElement | null>(null)
let map: L.Map | null = null
let markerLayer: L.LayerGroup | null = null

const reverseGeocode = async (lat: number, lng: number) => {
  try {
    const res = await fetch(`https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=${lat}&lon=${lng}&accept-language=id`)
    const data = await res.json()
    return data.display_name || 'Alamat tidak ditemukan'
  } catch (error) {
    console.error('Reverse geocoding error:', error)
    return 'Gagal mendeteksi alamat'
  }
}

onMounted(async () => {
  if (!mapContainer.value) return

  await store.fetchPoints()

  map = L.map(mapContainer.value, { zoomControl: false }).setView([-8.65, 115.2166], 12)
  L.control.zoom({ position: 'bottomleft' }).addTo(map)

  L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors',
    maxZoom: 19,
  }).addTo(map)

  markerLayer = L.layerGroup().addTo(map)

  map.on('click', async (e: L.LeafletMouseEvent) => {
    if (!store.isEditMode) return

    emit('map-clicked', { lat: e.latlng.lat, lng: e.latlng.lng })

    if (map) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent('<p class="font-bold text-sm text-primary m-0">⏳ Mendeteksi Alamat...</p>')
        .openOn(map)
    }

    const address = await reverseGeocode(e.latlng.lat, e.latlng.lng)

    if (store.activePoint && store.activePoint.latitude === e.latlng.lat && store.activePoint.longitude === e.latlng.lng) {
      store.activePoint = { ...store.activePoint, address }
    }

    if (map) {
      L.popup()
        .setLatLng(e.latlng)
        .setContent(`<div class="font-bold text-sm text-primary max-w-[200px] leading-tight"><p class="m-0 mb-1">📍 ${address}</p></div>`)
        .openOn(map)
    }
  })

  renderMarkers()

  setTimeout(() => map?.invalidateSize(), 100)
  setTimeout(() => map?.invalidateSize(), 500)
})

onUnmounted(() => {
  if (map) {
    map.remove()
    map = null
  }
})

const renderMarkers = () => {
  if (!map || !markerLayer) return
  markerLayer.clearLayers()

  const colors = [
    '#2D6A4F', '#0077B6', '#D62828', '#F4A261', '#9D4EDD',
    '#FFB703', '#E63946', '#219EBC', '#023047', '#FB8500',
    '#8ECAE6', '#606C38'
  ]

  filteredPoints.value.forEach((point) => {
    const typeObj = objectTypes.value.find(t => t.id === point.type_id)
    const typeName = typeObj ? typeObj.name : 'Unknown'
    const typeIconPath = typeObj ? typeObj.icon : 'M12 2v20M2 12h20'
    const color = colors[(point.type_id - 1) % colors.length]

    const customIcon = L.divIcon({
      className: 'custom-div-icon',
      html: `
        <div class="relative group">
          <div class="w-8 h-8 rounded-xl flex items-center justify-center text-white shadow-lg transition-all duration-300 group-hover:scale-110 group-hover:rotate-3 active:scale-95" style="background:${color}; border: 2px solid white;">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="w-4 h-4">
              <path d="${typeIconPath}" />
            </svg>
          </div>
          <div class="absolute -top-1 -right-1 w-3 h-3 bg-white rounded-full border-2 border-[${color}] shadow-sm animate-pulse"></div>
        </div>
      `,
      iconSize: [32, 32],
      iconAnchor: [16, 16],
      popupAnchor: [0, -24],
    })

    const marker = L.marker([point.latitude, point.longitude], { icon: customIcon }).addTo(markerLayer!)

    // Create a container for the Vue component
    const container = document.createElement('div')
    const vnode = h(MapPopupContent, { point, typeName })
    render(vnode, container)

    marker.bindPopup(container, {
      maxWidth: 300,
      className: 'premium-popup'
    })
  })
}

watch([filteredPoints, isEditMode, () => authStore.user], () => {
  renderMarkers()
})

watch(isEditMode, (newMode) => {
  if (mapContainer.value) {
    mapContainer.value.style.cursor = newMode ? 'crosshair' : 'grab'
  }
})

watch(isSidebarExpanded, () => {
  nextTick(() => {
    setTimeout(() => map?.invalidateSize(), 100)
    setTimeout(() => map?.invalidateSize(), 550)
  })
})

watch(() => store.flyToCoords, (newCoords) => {
  if (newCoords && map) {
    map.flyTo([newCoords.lat, newCoords.lng], 16)

    markerLayer?.eachLayer((layer: L.Layer) => {
      if (layer instanceof L.Marker) {
        const latLng = layer.getLatLng()
        if (latLng.lat === newCoords.lat && latLng.lng === newCoords.lng) {
          layer.openPopup()
        }
      }
    })
  }
})
</script>

<template>
  <div class="relative w-full h-full flex-1 min-h-0">
    <div ref="mapContainer" class="w-full h-full bg-slate-100 z-0"></div>
    <MapLegend />
  </div>
</template>

<style>
.leaflet-container {
  font-family: inherit;
}

.premium-popup .leaflet-popup-content-wrapper {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(8px);
  border: 1px solid rgba(0, 0, 0, 0.05);
  border-radius: 12px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.1);
  padding: 0;
}

.premium-popup .leaflet-popup-content {
  margin: 0;
}

.premium-popup .leaflet-popup-tip {
  background: white;
}
</style>
