import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMapUIStore = defineStore('mapUI', () => {
  const isSidebarExpanded = ref(false)

  // Map Interaction State
  const isEditMode = ref(false)
  const flyToCoords = ref<{ lat: number; lng: number } | null>(null)

  // Search & Filter State
  const searchQuery = ref('')
  const filterTypeId = ref<number | null>(null)
  const filterMyPoints = ref(false)

  // Toggle Actions
  const toggleSidebar = () => {
    isSidebarExpanded.value = !isSidebarExpanded.value
  }

  const toggleEditMode = () => {
    isEditMode.value = !isEditMode.value
  }

  const flyTo = (lat: number, lng: number) => {
    flyToCoords.value = { lat, lng }
  }

  return {
    isSidebarExpanded,
    isEditMode,
    flyToCoords,
    searchQuery,
    filterTypeId,
    filterMyPoints,
    toggleSidebar,
    toggleEditMode,
    flyTo,
  }
})
