<template>
    <div class="fixed pin z-50 bg-black-transparent flex">
        <div
            class="relative p-8 bg-white w-full max-w-sm mx-auto mb-auto mt-20 flex-col flex"
        >
            <div class="flex justify-between">
                <h2 class="font-normal">Request to Delete Data</h2>
                <XIcon
                    class="cursor-pointer text-grey hover:text-black"
                    @click="onClick"
                />
            </div>
            <div class="py-4">
                To request deletion of all of your personal data from
                [placeholder], please contact gdpr@placeholder.com
            </div>
        </div>
    </div>
</template>
<script>
import XIcon from 'vue-feather-icons/icons/XIcon';

export default {
    components: {
        XIcon
    },
    methods: {
        onClick() {
            this.$emit('close');
        }
    }
};
</script>
