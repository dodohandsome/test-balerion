<template>
  <div class="min-h-screen bg-[#1C1C1C] text-white">
    <div class="flex justify-end p-4">
      <div class="flex items-center space-x-4">
        <span>{{ authStore.user?.email }}</span>
        <button @click="authStore.logout" class="text-white hover:text-gray-300">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
              d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
          </svg>
        </button>
      </div>
    </div>
    <div class="max-w-6xl mx-auto py-8 px-4">
      <h1 class="text-center text-4xl font-bold mb-8">
        Memo Cards ({{ cards.length }})
      </h1>
      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6">
        <div v-for="card in cards" :key="card.id" class="bg-white rounded-lg p-6 text-black shadow-md">
          <div class="flex items-start space-x-4">
            <div class="flex flex-col">
              <span class="text-xl font-bold text-gray-700">
                {{ card.id }}
              </span>
              <span class="mt-2 px-3 py-1 rounded-full text-sm font-semibold" :class="card.tag === 'ADMIN'
                ? 'bg-red-100 text-red-800'
                : 'bg-blue-100 text-blue-800'">
                {{ card.tag }}
              </span>
            </div>
            <p class="flex-1 text-gray-800">
              {{ card.description }}
            </p>
          </div>
        </div>
        <button @click="showNewCardModal = true"
          class="bg-gray-700 rounded-lg p-6 flex items-center justify-center hover:bg-gray-600 shadow-md">
          <span class="text-4xl text-white">+</span>
        </button>
      </div>
    </div>
    <div v-if="showNewCardModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4">
      <div class="bg-white rounded-lg p-6 w-full max-w-md text-black">
        <h2 class="text-2xl font-bold mb-4">New Card</h2>
        <textarea v-model="newCardContent" class="w-full h-32 p-2 border rounded mb-4"
          placeholder="Enter card content..."></textarea>
        <div class="flex justify-end space-x-4">
          <button @click="showNewCardModal = false" class="px-4 py-2 text-gray-600 hover:text-gray-800">
            Cancel
          </button>
          <button @click="handleCreateCard" class="px-4 py-2 bg-black text-white rounded hover:bg-gray-800">
            Create
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useCardsStore } from '../stores/cards';
import { storeToRefs } from 'pinia';


const authStore = useAuthStore();
const cardsStore = useCardsStore();

const { cards } = storeToRefs(cardsStore);

const showNewCardModal = ref(false);
const newCardContent = ref('');
onMounted(() => {
  cardsStore.fetchCards();
});

async function handleCreateCard() {
  if (newCardContent.value.trim()) {
    try {
      await cardsStore.createCard(newCardContent.value.trim());
      newCardContent.value = '';
      showNewCardModal.value = false;
    } catch (error) {
      alert('Failed to create card');
    }
  }
}
// if (!authStore.isAuthenticated) {
//   router.push('/auth/login');
// }
</script>
