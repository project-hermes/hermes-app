<template>
    <main class="h-full w-full">
        <Modal v-if="isModalVisible" @close="isModalVisible = false" />
        <div
            class="w-full p-10 inline-flex items-center items-stretch border-b border-grey-lighter"
        >
            <h1 class="font-normal text-4xl">
                Account Settings
            </h1>
        </div>
        <div class="p-8">
            <div class="border border-grey-lighter mb-8">
                <div class="border-b border-grey-lighter">
                    <h3 class="font-normal p-4">Delete Data</h3>
                </div>
                <div class="w-full p-10 overflow-x-auto">
                    <button
                        class="bg-red p-3 text-white rounded"
                        @click="isModalVisible = true"
                    >
                        Request to Delete Data
                    </button>
                </div>
            </div>
        </div>
    </main>
</template>
<script>
import Modal from './Modal.vue';
export default {
    components: {
        Modal
    },
    data() {
        return {
            isModalVisible: false
        };
    }
};
</script>
