import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { Card } from '../api/cardsApi';
import { getCardsApi, createCardApi } from '../api/cardsApi';

export const useCardsStore = defineStore('cards', () => {
  const cards = ref<Card[]>([]);

  async function fetchCards() {
    try {
      const data = await getCardsApi();
      console.log(data)
      cards.value = data ? data : [];
      console.log(cards.value)
    } catch (error) {
      console.error('Failed to fetch cards:', error);
    }
  }

  async function createCard(description: string) {
    try {
      const newCard = await createCardApi({ description });
      // เพิ่ม card ใหม่เข้าไปใน state
      cards.value.push(newCard);
      console.log('Created new card:', newCard);
    } catch (error) {
      console.error('Failed to create card:', error);
      throw error;
    }
  }

  return {
    cards,
    createCard,
    fetchCards,
  };
});
