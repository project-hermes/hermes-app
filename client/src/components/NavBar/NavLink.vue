<template>
    <router-link
        :to="route"
        class="h-full px-1 uppercase no-underline text-xs flex items-center bg-blue hover:bg-blue-dark cursor-pointer sm:px-2 md:px-4 sm:text-sm md:mx-1"
        :class="{
            'text-white': isActive,
            'text-blue-light': !isActive
        }"
    >
        <slot />
    </router-link>
</template>
<script>
    export default {
        props: ['route']
    }
</script>
<style scoped>
.router-link-active {
    color: #fff;
}
</style>
