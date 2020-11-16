import axios from 'axios';
import { BasePoints } from './basePoints';

/*
  ダメージ一覧
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
          let d = this.ToDefendersResult(damages[i]);
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
    return new DefendersResult(d.target, d.damage_min, d.damage_max, d.rate_min, d.rate_max, d.determine_count);
  }
}

export class DefendersResult {
  readonly Target: string;
  readonly DamageMin: number;
  readonly DamageMax: number;
  readonly RateMin: number;
  readonly RateMax: number;
  readonly DetermineCount: number;

  constructor(target: string, damageMin: number, damageMax: number, rateMin: number, rateMax: number, determineCount: number) {
    this.Target = target;
    this.DamageMin = damageMin;
    this.DamageMax = damageMax;
    this.RateMax = rateMax;
    this.RateMin = rateMin;
    this.DetermineCount = determineCount;
  }

  Initialized(): boolean {
    return this.Target.length > 0;
  }

  static detault(): DefendersResult {
    return new DefendersResult('', 0, 0, 0, 0, 0);
  }
}

