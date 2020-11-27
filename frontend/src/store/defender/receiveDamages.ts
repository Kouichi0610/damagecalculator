import axios from 'axios'
import { TargetCondition } from '../target/targetCondition'

export class ReceiveDamages {
  receiveDamages(condition: TargetCondition, move: string): Promise<Result[]> {
    return new Promise((resolve, reject) => {
      axios.get('attacker_damages', {
        params: {
          Level: condition.level,
          BaseHP: condition.basePoints.hp(),
          BaseAttack: condition.basePoints.attack(),
          BaseDefense: condition.basePoints.defense(),
          BaseSpAttack: condition.basePoints.spAttack(),
          BaseSpDefense: condition.basePoints.spDefense(),
          BaseSpeed: condition.basePoints.speed(),
          Individuals: condition.individuals.type(),
          Name: condition.target,
          Ability: condition.ability,
          Nature: condition.nature.name,
          Item: condition.item.name,
          Condition: condition.condition,
          Weather: condition.weather.name,
          Field: condition.field.name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let damages = JSON.parse(json);
        let res: Result[] = []

        for (var i = 0; i < damages.length; i++) {
          let d = new Result(damages[i]);
          res.push(d);
        }
        resolve(res);
      })
      .catch((e) => {
        console.log('failed:' + e);
        let res: Result[] = []
        reject(res);
      });
    });
  }
}

export class Result {
  readonly target: string = '';
  readonly move: string = '';
  readonly damageMin: number = 0;
  readonly damageMax: number = 0;
  readonly rateMin: number = 0;
  readonly rateMax: number = 0;
  readonly determineCount: number = 0;

  toString(): string {
    return this.target
     + ' ' + this.move
     + ' ダメージ' + this.damageMin + '～' + this.damageMax
     + ' ' + this.rateMin + '%～' + this.rateMax + '%'
     + ' 確定数' + this.determineCount;
  }

  constructor(d: any) {
    if (d == null) {
      return;
    }
    this.target = d.target;
    this.move = d.move;
    this.damageMin = d.damage_min;
    this.damageMax = d.damage_max;
    this.rateMax = d.rate_max;
    this.rateMin = d.rate_min;
    this.determineCount = d.determine_count;
  }
}
