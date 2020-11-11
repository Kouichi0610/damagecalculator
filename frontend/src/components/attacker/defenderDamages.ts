import axios from 'axios';

/*
  ダメージ一覧を取得
*/
export class DefenderDamages {
  private level: number
  private attacker: string;
  private individuals: string;
  private basePoints: number[];
  private nature: string;
  private ability: string;
  private move: string;
  private item: string;
  private condition: string;
  private weather: string;
  private field: string;

  constructor(attacker: string, level: number, individuals: string, basePoints: number[], nature: string, ability: string, move: string, item: string, condition: string, weather: string, field: string) {
     
    this.attacker = attacker;
    this.level = level;
    this.individuals = individuals;
    this.basePoints = basePoints;
    this.ability = ability;
    this.nature = nature;
    this.move = move;
    this.item = item;
    this.condition = condition;
    this.weather = weather;
    this.field = field;
  }
  // 試作
  getSample() {
    axios.get('defender_damages', {
      params: {
        Level: this.level,
        BaseHP: this.basePoints[0],
        BaseAttack: this.basePoints[1],
        BaseDefense: this.basePoints[2],
        BaseSpAttack: this.basePoints[3],
        BaseSpDefense: this.basePoints[4],
        BaseSpeed: this.basePoints[5],
        Individuals: this.individuals,
        Name: this.attacker,
        Move: this.move,
        Ability: this.ability,
        Nature: this.nature,
        Item: this.item,
        Condition: this.condition,
        Weather: this.weather,
        Field: this.field,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      //let descriptions = JSON.parse(json);
      console.log('' + json);
      console.log('success.');
      return true;
    })
    .catch((e) => {
      console.log('failed:' + e);
      return false;
    });
  }
}

export interface DefendersResult {

}

/*
[{"target":"カビゴン ずぶとい ","damage_min":107,"damage_max":127,"rate_min":40.074906367041194,"rate_max":47.565543071161045,"determine_count":3},
{"target":"ハピナス おだやか ","damage_min":68,"damage_max":81,"rate_min":18.784530386740332,"rate_max":22.375690607734807,"determine_count":6},{"target":"ドヒドイデ わんぱく ","damage_min":86,"damage_max":102,"rate_min":54.77707006369427,"rate_max":64.96815286624204,"determine_count":2},{"target":"ウォッシュロトム ひかえめ ","damage_min":109,"damage_max":129,"rate_min":69.42675159235668,"rate_max":82.16560509554141,"determine_count":2},{"target":"フシギバナ ひかえめ ","damage_min":115,"damage_max":136,"rate_min":61.49732620320856,"rate_max":72.72727272727273,"determine_count":2},{"target":"ユキノオー いじっぱり ","damage_min":131,"damage_max":155,"rate_min":66.49746192893402,"rate_max":78.68020304568529,"determine_count":2},{"target":"ジャラランガ いじっぱり ","damage_min":111,"damage_max":131,"rate_min":60.98901098901099,"rate_max":71.97802197802197,"determine_count":2},{"target":"マタドガス わんぱく ","damage_min":153,"damage_max":181,"rate_min":88.95348837209302,"rate_max":105.23255813953489,"determine_count":2},{"target":"カバルドン わんぱく ","damage_min":150,"damage_max":177,"rate_min":69.76744186046511,"rate_max":82.32558139534883,"determine_count":2},{"target":"マンタイン おだやか ","damage_min":67,"damage_max":79,"rate_min":34.89583333333333,"rate_max":41.14583333333333,"determine_count":3},{"target":"クレセリア ひかえめ ","damage_min":93,"damage_max":110,"rate_min":40.969162995594715,"rate_max":48.458149779735685,"determine_count":3},{"target":"ストライク ようき ","damage_min":140,"damage_max":165,"rate_min":95.8904109589041,"rate_max":113.013698630137,"determine_count":2},{"target":"ガブリアス ようき ","damage_min":133,"damage_max":157,"rate_min":72.28260869565217,"rate_max":85.32608695652173,"determine_count":2},{"target":"ニンフィア ひかえめ ","damage_min":93,"damage_max":110,"rate_min":46.03960396039604,"rate_max":54.45544554455446,"determine_count":3}]
*/