import axios from 'axios';

/*
  ダメージ一覧試作
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
    return new DefendersResult(d.target, d.damage_min, d.damage_max, d.rate_min, d.rate_max, d.determine_count);
  }
}

export class DefendersResult {
  Target: string;
  DamageMin: number;
  DamageMax: number;
  RateMin: number;
  RateMax: number;
  DetermineCount: number;

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

  static Defaults(count: number): DefendersResult[] {
    let res: DefendersResult[] = [];
    for (var i = 0; i < count; i++) {
      res.push(this.Default());
    }
    return res;
  }

  static Default(): DefendersResult {
    return new DefendersResult('', 0, 0, 0, 0, 0);
  }
}

