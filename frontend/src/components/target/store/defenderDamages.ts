import axios from 'axios';
import { BasePoints } from './basePoints';
import { DamageRateInfo } from './damageRateInfo';

/*
  攻撃調整
  仮想敵へのダメージ一覧
*/
export class DefenderDamages {
  readonly attacker: string;
  private level: number
  private individuals: string;
  private basePoints: BasePoints;
  private nature: string;
  private ability: string;
  private item: string;
  private condition: string;
  private weather: string;
  private field: string;

  constructor(attacker: string, level: number, individuals: string, basePoints: BasePoints, nature: string, ability: string, item: string, condition: string, weather: string, field: string) {
    this.attacker = attacker;
    this.level = level;
    this.individuals = individuals;
    this.basePoints = basePoints;
    this.ability = ability;
    this.nature = nature;
    this.item = item;
    this.condition = condition;
    this.weather = weather;
    this.field = field;
  }

  defenderDamages(move: string): Promise<DefendersResult[]> {
    return new Promise((resolve, reject) => {
      axios.get('defender_damages', {
        params: {
          Level: this.level,
          BaseHP: this.basePoints.hp,
          BaseAttack: this.basePoints.attack,
          BaseDefense: this.basePoints.defense,
          BaseSpAttack: this.basePoints.spAttack,
          BaseSpDefense: this.basePoints.spDefense,
          BaseSpeed: this.basePoints.speed,
          Individuals: this.individuals,
          Name: this.attacker,
          Move: move,
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
          let d = new DefendersResult(damages[i]);
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
}

export class DefendersResult implements DamageRateInfo{
  readonly Target: string = '';
  readonly DamageMin: number = 0;
  readonly DamageMax: number = 0;
  readonly RateMin: number = 0;
  readonly RateMax: number = 0;
  readonly DetermineCount: number = 0;

  toString(): string {
    return this.Target
     + ' ダメージ' + this.DamageMin + '～' + this.DamageMax
     + ' ' + this.RateMin + '%～' + this.RateMax + '%'
     + ' 確定数' + this.DetermineCount;
  }

  constructor(d: any) {
    if (d == null) {
      return;
    }
    this.Target = d.target;
    this.DamageMin = d.damage_min;
    this.DamageMax = d.damage_max;
    this.RateMax = d.rate_max;
    this.RateMin = d.rate_min;
    this.DetermineCount = d.determine_count;
  }

  Initialized(): boolean {
    return this.Target.length > 0;
  }

  static detault(): DefendersResult {
    return new DefendersResult(null);
  }
}

