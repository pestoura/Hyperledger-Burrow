//Code generated by solts. DO NOT EDIT.
import { Address, CancelStreamSignal, ContractCodec, Event, Keccak, listenerFor } from '../../index';
interface Provider {
  deploy(
    data: string | Uint8Array,
    contractMeta?: {
      abi: string;
      codeHash: Uint8Array;
    }[],
  ): Promise<Address>;
  call(data: string | Uint8Array, address: string): Promise<Uint8Array | undefined>;
  callSim(data: string | Uint8Array, address: string): Promise<Uint8Array | undefined>;
  listen(
    signatures: string[],
    address: string,
    callback: (err?: Error, event?: Event) => CancelStreamSignal | void,
    start?: 'first' | 'latest' | 'stream' | number,
    end?: 'first' | 'latest' | 'stream' | number,
  ): unknown;
  contractCodec(contractABI: string): ContractCodec;
}
export type Caller = typeof defaultCall;
export async function defaultCall<Output>(
  client: Provider,
  addr: string,
  data: Uint8Array,
  isSim: boolean,
  callback: (returnData: Uint8Array | undefined) => Output,
): Promise<Output> {
  const returnData = await (isSim ? client.callSim(data, addr) : client.call(data, addr));
  return callback(returnData);
}
export namespace Eventer {
  export const contractName = 'Eventer';
  export const abi =
    '[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"eventId","type":"bytes32"},{"indexed":true,"internalType":"bytes32","name":"intervalId","type":"bytes32"},{"indexed":false,"internalType":"address","name":"eventAddress","type":"address"},{"indexed":false,"internalType":"string","name":"namespace","type":"string"},{"indexed":false,"internalType":"string","name":"name","type":"string"},{"indexed":false,"internalType":"address","name":"controller","type":"address"},{"indexed":false,"internalType":"uint256","name":"threshold","type":"uint256"},{"indexed":false,"internalType":"string","name":"metadata","type":"string"}],"name":"Init","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"int256","name":"timestamp","type":"int256"},{"indexed":false,"internalType":"string","name":"place","type":"string"},{"indexed":false,"internalType":"string","name":"postalAddress","type":"string"}],"name":"MonoRampage","type":"event"},{"constant":false,"inputs":[],"name":"announce","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]';
  export const bytecode =
    '6080604052348015600f57600080fd5b506102b58061001f6000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638f970df014610030575b600080fd5b61003861003a565b005b607b7f6a35688e78094e86ac7dd4593423fa89415105dc68a0766b27106861ef4102146040518080602001806020018381038352600d8152602001807f53616e74612045756c61726961000000000000000000000000000000000000008152506020018381038252600a8152602001807f53616e74204a75616d65000000000000000000000000000000000000000000008152506020019250505060405180910390a27f696e74657276616c3200000000000000000000000000000000000000000000007f6576656e743100000000000000000000000000000000000000000000000000007f5f20df97ee573ab8b43581cf3ff905f3507ad2329b7efe6f92e802b4fad031c17359c99d4ebf520619ee7f806f11d90a9cac02ce06336004604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200180602001848103845260068152602001807f64696e696e670000000000000000000000000000000000000000000000000000815250602001848103835260098152602001807f627265616b666173740000000000000000000000000000000000000000000000815250602001848103825260178152602001807f6261636f6e2c6265616e732c656767732c746f6d61746f000000000000000000815250602001965050505050505060405180910390a356fea265627a7a72315820a6cd65593c9a8f5ed08f75519d9dd5664bcdf3a4bb67c8f507e1a69ed348f04f64736f6c63430005110032';
  export const deployedBytecode =
    '608060405234801561001057600080fd5b506004361061002b5760003560e01c80638f970df014610030575b600080fd5b61003861003a565b005b607b7f6a35688e78094e86ac7dd4593423fa89415105dc68a0766b27106861ef4102146040518080602001806020018381038352600d8152602001807f53616e74612045756c61726961000000000000000000000000000000000000008152506020018381038252600a8152602001807f53616e74204a75616d65000000000000000000000000000000000000000000008152506020019250505060405180910390a27f696e74657276616c3200000000000000000000000000000000000000000000007f6576656e743100000000000000000000000000000000000000000000000000007f5f20df97ee573ab8b43581cf3ff905f3507ad2329b7efe6f92e802b4fad031c17359c99d4ebf520619ee7f806f11d90a9cac02ce06336004604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001806020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200180602001848103845260068152602001807f64696e696e670000000000000000000000000000000000000000000000000000815250602001848103835260098152602001807f627265616b666173740000000000000000000000000000000000000000000000815250602001848103825260178152602001807f6261636f6e2c6265616e732c656767732c746f6d61746f000000000000000000815250602001965050505050505060405180910390a356fea265627a7a72315820a6cd65593c9a8f5ed08f75519d9dd5664bcdf3a4bb67c8f507e1a69ed348f04f64736f6c63430005110032';
  export function deploy({
    client,
    withContractMeta,
  }: {
    client: Provider;
    withContractMeta?: boolean;
  }): Promise<string> {
    const codec = client.contractCodec(abi);
    const data = Buffer.concat([Buffer.from(bytecode, 'hex'), codec.encodeDeploy()]);
    return client.deploy(
      data,
      withContractMeta
        ? [{ abi: Eventer.abi, codeHash: new Keccak(256).update(Eventer.deployedBytecode, 'hex').digest('binary') }]
        : undefined,
    );
  }
  export async function deployContract(deps: { client: Provider; withContractMeta?: boolean }): Promise<Contract> {
    const address = await deploy(deps);
    return contract(deps.client, address);
  }
  type EventRegistry = typeof events;
  export type EventName = keyof EventRegistry;
  export type TaggedPayload<T extends EventName> = ReturnType<EventRegistry[T]['tagged']> & {
    event: Event;
  };
  export type SolidityEvent<T extends EventName> = TaggedPayload<T>['payload'];
  export type TypedListener<T extends EventName> = (
    callback: (err?: Error, event?: TaggedPayload<T>) => void,
    start?: 'first' | 'latest' | 'stream' | number,
    end?: 'first' | 'latest' | 'stream' | number,
  ) => unknown;
  const events = {
    Init: {
      signature: '5F20DF97EE573AB8B43581CF3FF905F3507AD2329B7EFE6F92E802B4FAD031C1',
      tagged: (
        eventId: Buffer,
        intervalId: Buffer,
        eventAddress: string,
        namespace: string,
        name: string,
        controller: string,
        threshold: number,
        metadata: string,
      ) =>
        ({
          name: 'Init',
          payload: {
            eventId: eventId,
            intervalId: intervalId,
            eventAddress: eventAddress,
            namespace: namespace,
            name: name,
            controller: controller,
            threshold: threshold,
            metadata: metadata,
          } as const,
        } as const),
    } as const,
    MonoRampage: {
      signature: '6A35688E78094E86AC7DD4593423FA89415105DC68A0766B27106861EF410214',
      tagged: (timestamp: number, place: string, postalAddress: string) =>
        ({
          name: 'MonoRampage',
          payload: { timestamp: timestamp, place: place, postalAddress: postalAddress } as const,
        } as const),
    } as const,
  } as const;
  export type Contract = ReturnType<typeof contract>;
  export const contract = (client: Provider, address: string) =>
    ({
      name: 'Eventer',
      address,
      functions: {
        announce(call = defaultCall): Promise<void> {
          const data = encode(client).announce();
          return call<void>(client, address, data, false, (data: Uint8Array | undefined) => {
            return decode(client, data).announce();
          });
        },
      } as const,
      listeners: {
        Init(
          callback: (
            err?: Error,
            event?: {
              eventId: Buffer;
              intervalId: Buffer;
              eventAddress: string;
              namespace: string;
              name: string;
              controller: string;
              threshold: number;
              metadata: string;
            },
          ) => CancelStreamSignal | void,
          start?: 'first' | 'latest' | 'stream' | number,
          end?: 'first' | 'latest' | 'stream' | number,
        ): unknown {
          return client.listen(
            ['5F20DF97EE573AB8B43581CF3FF905F3507AD2329B7EFE6F92E802B4FAD031C1'],
            address,
            (err?: Error, event?: Event) => {
              if (err) {
                return callback(err);
              }
              return callback(undefined, decode(client, event?.log.data, event?.log.topics).Init());
            },
            start,
            end,
          );
        },
        MonoRampage(
          callback: (
            err?: Error,
            event?: {
              timestamp: number;
              place: string;
              postalAddress: string;
            },
          ) => CancelStreamSignal | void,
          start?: 'first' | 'latest' | 'stream' | number,
          end?: 'first' | 'latest' | 'stream' | number,
        ): unknown {
          return client.listen(
            ['6A35688E78094E86AC7DD4593423FA89415105DC68A0766B27106861EF410214'],
            address,
            (err?: Error, event?: Event) => {
              if (err) {
                return callback(err);
              }
              return callback(undefined, decode(client, event?.log.data, event?.log.topics).MonoRampage());
            },
            start,
            end,
          );
        },
      } as const,
      listenerFor: <T extends EventName>(eventNames: T[]): TypedListener<T> =>
        (listenerFor(client, address, events, decode, eventNames) as unknown) as TypedListener<T>,
      listener: listenerFor(
        client,
        address,
        events,
        decode,
        Object.keys(events) as EventName[],
      ) as TypedListener<EventName>,
    } as const);
  export const encode = (client: Provider) => {
    const codec = client.contractCodec(abi);
    return {
      announce: () => {
        return codec.encodeFunctionData('8F970DF0');
      },
    };
  };
  export const decode = (client: Provider, data: Uint8Array | undefined, topics: Uint8Array[] = []) => {
    const codec = client.contractCodec(abi);
    return {
      Init: (): {
        eventId: Buffer;
        intervalId: Buffer;
        eventAddress: string;
        namespace: string;
        name: string;
        controller: string;
        threshold: number;
        metadata: string;
      } => {
        const [
          eventId,
          intervalId,
          eventAddress,
          namespace,
          name,
          controller,
          threshold,
          metadata,
        ] = codec.decodeEventLog('5F20DF97EE573AB8B43581CF3FF905F3507AD2329B7EFE6F92E802B4FAD031C1', data, topics);
        return {
          eventId: eventId,
          intervalId: intervalId,
          eventAddress: eventAddress,
          namespace: namespace,
          name: name,
          controller: controller,
          threshold: threshold,
          metadata: metadata,
        };
      },
      MonoRampage: (): {
        timestamp: number;
        place: string;
        postalAddress: string;
      } => {
        const [timestamp, place, postalAddress] = codec.decodeEventLog(
          '6A35688E78094E86AC7DD4593423FA89415105DC68A0766B27106861EF410214',
          data,
          topics,
        );
        return { timestamp: timestamp, place: place, postalAddress: postalAddress };
      },
      announce: (): void => {
        return;
      },
    };
  };
}
