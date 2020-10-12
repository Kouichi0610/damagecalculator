<template>
  <div class="hello">
    <input @click="getSample" type="button" value="GetSample">
    <input @click="postSample" type="button" value="PostSample">
    <input @click="getNames" type="button" value="Names">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
/*
class postData {
  constructor(x, y) {
    this.x = x;
    this.y = y;
  }
  toURLSearchParams() {
    let res = new URLSearchParams();
    res.append('x', this.x);
    res.append('y', this.y);
    return res;
  }
    //JSON形式の場合
    //axios.post('post_sample', {
    //  x: this.x,
    //  y: this.y,
    //})
}
*/

export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  methods: {
    getNames() {
      this.axios.get('get_names')
      .then((response) => {
        alert('success : ' + response.data)
      })
      .catch((e) => {
        alert('failed ' + e)
      })
    },
    postSample() {
      //let data = new postData(11, 97);
      const data = { x: '11', y: '97' };
      this.axios.post('post_sample', data)
      .then((response) => {
        alert('success : ' + response.data)
      })
      .catch((e) => {
        alert('failed ' + e)
      })
    },
    getSample() {
      // クエリパラメータ付き
      // URLに直接付けてもいい'get_sample?Name="ゴローニャ"'
      this.axios.get('get_sample', {
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

        alert('success? : ' + msg)
      })
      .catch((e) => {
        alert('failed ' + e)
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
