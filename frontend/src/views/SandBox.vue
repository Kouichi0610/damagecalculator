<template>
  <div class="sandbox">
    <h1>SandBox</h1>
    <stats-editor></stats-editor>
    名前 {{ pokename }} Count {{ count }}<br>
    <damage-result :min="rateMin" :max="rateMax"></damage-result>
    <names v-bind:tag="pokemon" v-bind:name="name" v-bind:candidates="candidates"></names>
    <rank></rank>
  </div>
</template>

<script>
import StatsEditor from '../components/statsEditor'

import Names from '../components/names'
import Rank from '../components/rank'
import DamageResult from '../components/damageResult'
import store from "../store/index"

export default {
  name: 'SandBox',
  components: {
    'stats-editor': StatsEditor,
    'names': Names,
    'rank': Rank,
    'damage-result': DamageResult,
  },
  data: function() {
    return {
      rateMin: '16.7',
      rateMax: '25.4',
      pokemon: 'ポケモン',
      name: 'ピカチュウ',
      candidates: ['アーボック', 'フシギバナ', 'ギャラドス', 'リザードン', 'カメックス', 'ピカチュウ', 'カイリュー', 'ドサイドン', 'ミュウ', 'ホルード', 'ランクルス', 'エルレイド', 'バタフリー'],
    }
  },
  created: function() {
    console.log('Created');
    console.log('pokename ' + store.state.pokename);

    store.commit('stats/nature/setNature', 'ずぶとい');
    store.commit('stats/individuals/toWeakest', true);
    store.commit('stats/individuals/toMax');
    store.commit('stats/individuals/toSlowest', true);
    store.commit('counter/increment');
    store.dispatch('counter/addAsync', {
      amount: 1000
    });
  },
  computed: {
    count() {
      return store.state.counter.count;
    },
    pokename() {
      return store.getters.statsString;
      //return store.getters.myname + store.state.stats.sample + store.state.stats.individuals.toString();
      //return store.state.pokename;
    }
  },
  watch: {
    name: function(next, prev) {
      console.log('Prev' + prev + ' Next:' + next);
    }

  }
}
</script>

<style scoped>
</style>