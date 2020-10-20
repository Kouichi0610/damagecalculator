<template>
    <div class="sandboxts">
        SandBox from TypeScript.
        <Todo></Todo>
        <user-detail></user-detail>
        <StatsEditor></StatsEditor>
    </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

import StatsEditor from '../components/stats/statsEditor.vue'
import UserDetail from '../components/userDetail.vue'
import Todo from '../components/todo.vue'

@Component({
      components: {
          StatsEditor,
          UserDetail,
          Todo,
    },
})
export default class SandBoxTs extends Vue {
    species: number[] = [100, 25, 180, 35, 120, 60];
}
</script>

<style scoped>
</style>
