<template>
    <div class="sandboxts">
        <Target :targetName="name"></Target>
        <TargetSelect @selectPokemon="onNameChanged"></TargetSelect>
        SandBox from TypeScript.(1.1)
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import TargetSelect from '../components/targetSelect/targetSelect.vue'
import Target from '../components/target/target.vue'

@Component({
      components: {
          TargetSelect,
          Target,
    },
})
export default class SandBoxTs extends Vue {
    species: number[] = [100, 25, 180, 35, 120, 60];

    private name: string = ''

    onNameChanged(name: string) {
        this.name = name;
    }
}
</script>

<style scoped>
</style>
