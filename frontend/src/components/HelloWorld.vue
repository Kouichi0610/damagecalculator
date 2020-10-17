<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <input @click="getSample" type="button" value="GetSample">
    <input @click="postSample" type="button" value="PostSample">
    <input @click="filteredList" type="button" value="リスト取得">
    <p>
      For a guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
    </p>
    <h3>Installed CLI Plugins</h3>
    <ul>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-typescript" target="_blank" rel="noopener">typescript</a></li>
    </ul>
    <h3>Essential Links</h3>
    <ul>
      <li><a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a></li>
      <li><a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a></li>
      <li><a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a></li>
      <li><a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a></li>
    </ul>
    <h3>Ecosystem</h3>
    <ul>
      <li><a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a></li>
      <li><a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a></li>
      <li><a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank" rel="noopener">vue-devtools</a></li>
      <li><a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a></li>
      <li><a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a></li>
    </ul>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import axios from 'axios'

// TODO:使っているのにエラーが出る、わけわからない
//import {Species} from '../store/targets/types'

class SpeciesImpl implements Species {
    name: string;
    types: string;
    hp: number;
    attack: number;
    defense: number;
    spAttack: number;
    spDefense: number;
    speed: number;

    toString():string{
      return '' + this.name;
    }

    constructor (d: any) {
      this.name = d.name;
      this.types = d.types;
      this.hp = d.hp;
      this.attack = d.attack;
      this.defense = d.defense;
      this.spAttack = d.sp_attack;
      this.spDefense = d.sp_defense;
      this.speed = d.speed;
    }
}

/*
  クエリ指定方法1
  axios.get('filtered_list?Types=はがね&Total=540')
  クエリ指定方法2
  axios.get('filtered_list', {
    params: {
      Types: 'ほのお',
      Total: '530',
    }
  })
*/

@Component
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;
  filteredList(): void {
      //axios.get('filtered_list?Types=はがね&Total=540')
      axios.get('filtered_list', {
        params: {
          Types: 'ドラゴン',
          Total: '530',
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        console.log('json:' + json)

        let sp = JSON.parse(json);
        let species: Species[] = new Array();
        for (var i = 0; i < sp.length; i++) {
          let s = new SpeciesImpl(sp[i]);
          species.push(s);
        }
        for (i = 0; i < species.length; i++) {
          let s = species[i]
          console.log('Species:' + i + ':' + s.name + ' HP:' + s.hp + ' Attack:' + s.attack 
          + ' Defense:' + s.defense + ' SpAttack:' + s.spAttack + ' SpDefense:' + s.spDefense + ' Speed:' + s.speed);


        }
      })
      .catch((e) => {
        alert('failed ' + e)
      })
  }
  postSample(): void {
      const data = { x: '11', y: '97' };
      axios.post('post_sample', data)
      .then((response) => {
        alert('success : ' + response.data)
      })
      .catch((e) => {
        alert('failed ' + e)
      })
  }
  getSample(): void {
      // クエリパラメータ付き
      // URLに直接付けてもいい'get_sample?Name="ゴローニャ"'
      axios.get('get_sample', {
        params: {
          Name: 'ゴローニャ',
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        console.log('json:' + json);
        let res = JSON.parse(json);
        var msg = '';
        msg += res[0].name;
        msg += ' HP:' + res[0].hp;
        msg += ' AT:' + res[0].attack;
        msg += ' DF:' + res[0].defense;
        msg += ' SA:' + res[0].sp_attack;
        msg += ' SD:' + res[0].sp_defense;
        msg += ' SP:' + res[0].speed;
        alert('success:' + msg);
      })
      .catch((e) => {
        alert('failed ' + e)
      })
  }


}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
