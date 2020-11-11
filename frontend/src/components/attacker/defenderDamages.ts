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

  defenderDamages(): Promise<DefendersResult[]> {
    return new Promise((resolve, reject) => {
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
        let damages = JSON.parse(json);
        let res: DefendersResult[] = [];
  
        for (var i = 0; i < damages.length; i++) {
          let d = this.ToDefendersResult(damages[i]);
          console.log('' + d.Target + 'Dmg:' + d.DamageMin + '-' + d.DamageMax + 'Rate:' + d.RateMin + '-' + d.RateMax + ' 確定数' + d.DetermineCount);
          res.push(d);
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('error:' + e);
        let res: DefendersResult[] = [];
        reject(res);
      });
    });
  }

  private ToDefendersResult(d: any): DefendersResult {
    return {
      Target:d.target,
      DamageMin: d.damage_min,
      DamageMax: d.damage_max,
      RateMin: d.rate_min,
      RateMax: d.rate_max,
      DetermineCount: d.determine_count,
    }
  }
}

export interface DefendersResult {
  Target: string;
  DamageMin: number;
  DamageMax: number;
  RateMin: number;
  RateMax: number;
  DetermineCount: number;
}
