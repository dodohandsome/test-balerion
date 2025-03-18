import { api } from './axiosInstance';

export interface Card {
    id: string;
    tag: string;
    description: string;
    createdAt: string; // หรือใช้ Date แล้วแต่การ parse ในฝั่ง client
}

// ฟังก์ชันสำหรับดึงข้อมูลการ์ด
export async function getCardsApi(): Promise<Card[]> {
    const response = await api.get<Card[]>('/cards');
    return response.data;
}

export async function createCardApi(cardData: { description: string }): Promise<Card> {
    const response = await api.post<Card>('/cards', cardData);
    return response.data;
}
