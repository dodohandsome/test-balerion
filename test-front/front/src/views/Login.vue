<template>
  <!-- เต็มจอด้วย w-screen h-screen และตั้งสีพื้นหลังเป็นเทาเข้ม (#1C1C1C) -->
  <div class="w-screen h-screen bg-[#000000] flex items-center justify-center px-4">
    <!-- max-w-md เพื่อจำกัดความกว้างของฟอร์ม -->
    <div class="w-full max-w-md">
      <!-- หัวข้อใหญ่ สีขาว -->
      <h1 class="text-white text-2xl font-bold mb-8 text-left">
        เข้าสู่ระบบ
      </h1>

      <form @submit.prevent="handleLogin" class="space-y-6 text-left">
        <!-- บัญชีพนักงาน -->
        <div>
          <label class="block text-gray-300 text-sm mb-2">บัญชีพนักงาน</label>
          <input v-model="username" type="text"
            class="w-full px-3 py-2 bg-[#2F2F2F] rounded text-white placeholder-gray-400 outline-none focus:ring-2 focus:ring-gray-600"
            placeholder="A0001" />
        </div>

        <!-- รหัสผ่าน -->
        <div>
          <label class="block text-gray-300 text-sm mb-2">รหัสผ่าน</label>
          <input v-model="password" type="password"
            class="w-full px-3 py-2 bg-[#2F2F2F] rounded text-white placeholder-gray-400 outline-none focus:ring-2 focus:ring-gray-600"
            placeholder="------" />
        </div>

        <!-- ปุ่มค้นหา (พื้นหลังขาว ตัวหนังสือดำ) -->
        <button type="submit" class="w-full bg-white text-black py-2 rounded hover:bg-gray-200">
          ค้นหา
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';


const router = useRouter();
const authStore = useAuthStore();

const username = ref('');
const password = ref('');

async function handleLogin() {
  try {
    // เรียก login ผ่าน store
    await authStore.login(username.value, password.value);
    router.push('/');
  } catch (error) {
    alert('Invalid credentials');
  }
}
</script>
