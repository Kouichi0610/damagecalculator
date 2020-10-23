<template>
    <div class="TypesButton">
        <b-button class="TypeButton" v-on:click="onButton">{{ type }}</b-button>
    </div>
</template>

<script lang="ts">
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class TypesButton extends Vue {
    @Prop() private type!: string;

    onButton() {
        this.$emit('clicked', this.type);
    }
}
</script>

<style scoped>
.TypeButton {
    width: 108px;
}
.TypesButton {
}
</style>
