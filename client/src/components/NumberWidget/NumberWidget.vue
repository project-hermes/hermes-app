<template>
    <div>
        <label
            class="text-grey-darker text-sm leading-tight whitespace-no-wrap"
            >{{ label }}</label
        >
        <div
            v-if="value"
            :class="'text-' + color"
      class="text-2xl whitespace-no-wrap"
        >
            {{ value }}
        </div>
        <div v-else>---</div>
    </div>
</template>
<script>
export default {
    props: {
        label: {
            type: String,
            required: true
        },
        value: {
            type: [String, Number]
        },
        color: {
            type: String,
            default: 'blue'
        }
    }
};
</script>
