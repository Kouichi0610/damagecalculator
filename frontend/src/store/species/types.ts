import axios from 'axios'

export interface SpeciesState {
  species: ISpecies;
}

export class PokeData {
  readonly name: string = '';
  readonly types: string[] = [''];
  readonly weight: number = 0;
  readonly species: ISpecies = defaultSpecies();
  readonly abilities: string[] = [];

  hasTarget(): boolean {
    return this.name.length > 0;
  }

  static default(): PokeData {
    return new PokeData(null);
  }
  constructor(data: any) {
    if (data == null) return;
    this.name = data.Name;
    this.abilities = data.Abilities;
    this.types = data.Types;
    this.weight = data.Weight;
    let sp: number[] = data.Species;
    this.species = new Species(sp[0], sp[1], sp[2], sp[3], sp[4], sp[5]);
  }
  toString(): string {
    if (!this.hasTarget()) return 'no Target.';
    return '' + this.name
    + ' Types:' + this.types
    + ' Weight:' + this.weight + 'kg'
    + ' species:' + speciesToString(this.species)
    + ' abilities:' + this.abilities;
  }
}

export class PokeDataLoader {
  public load(name: string): Promise<PokeData> {
    return new Promise((resolve, reject) => {
      axios.get('get_species', {
        params: {
          Name: name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let sp = JSON.parse(json);
        let d = new PokeData(sp);
        resolve(d);
      })
      .catch((e) => {
        console.log('failed:' + e);
        reject(PokeData.default());
      });
    });
  }
}

function speciesToString(s: ISpecies): string {
  return 'HP:' + s.hp
  + ' 攻撃:' + s.attack
  + ' 防御:' + s.defense
  + ' 特攻:' + s.spAttack
  + ' 特防:' + s.spDefense
  + ' 素早さ:' + s.speed;
}

// 種族値
export interface ISpecies {
  readonly hp: number;
  readonly attack: number;
  readonly defense: number;
  readonly spAttack: number;
  readonly spDefense: number;
  readonly speed: number;
}

export function defaultSpecies(): ISpecies {
  return new Species(0,0,0,0,0,0);
}

class Species implements ISpecies {
  readonly hp: number = 0;
  readonly attack: number = 0;
  readonly defense: number = 0;
  readonly spAttack: number = 0;
  readonly spDefense: number = 0;
  readonly speed: number = 0;

  constructor(hp: number, attack: number, defense: number, spAttack: number, spDefense: number, speed: number) {
    this.hp = hp;
    this.attack = attack;
    this.defense = defense;
    this.spAttack = spAttack;
    this.spDefense = spDefense;
    this.speed = speed;
  }
}
