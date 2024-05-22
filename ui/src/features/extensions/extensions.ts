import { IconDefinition } from '@fortawesome/free-solid-svg-icons';

import { ModalComponentProps } from '../common/modal/modal-context';

interface GenericExtensionProps {}
type ModalExtensionProps = GenericExtensionProps & ModalComponentProps;

export interface CreateStageProps extends GenericExtensionProps {}
export interface ProjectTabProps extends GenericExtensionProps {}
export interface SystemTabProps extends GenericExtensionProps {}
export interface FreightTabExtensionProps extends GenericExtensionProps {}
export interface ProjectActionProps extends ModalExtensionProps {}

export enum ExtensionKind {
  CreateStage = 'CreateStage',
  ProjectTab = 'ProjectTab',
  SystemTab = 'SystemTab',
  FreightTab = 'FreightTab',
  ProjectAction = 'ProjectAction'
}

type TypeMapping = {
  [ExtensionKind.CreateStage]: CreateStageProps;
  [ExtensionKind.ProjectTab]: ProjectTabProps;
  [ExtensionKind.SystemTab]: SystemTabProps;
  [ExtensionKind.ProjectAction]: ProjectActionProps;
  [ExtensionKind.FreightTab]: FreightTabExtensionProps;
};

type PairedType<T extends ExtensionKind> = TypeMapping[T];

export interface Extension<T extends ExtensionKind> {
  component: (props: PairedType<T>) => React.ReactNode;
  version: string;
  kind: T;
  name: string;
  label?: string;
  icon?: IconDefinition;
}

type ExtensionsOfKind<T extends ExtensionKind> = Record<string, Extension<T>>;

type MappedRecord<K extends ExtensionKind> = {
  [P in K]: ExtensionsOfKind<P>;
};

const ExtensionsStore = {} as MappedRecord<ExtensionKind>;

export function useExtensions<T extends ExtensionKind>(kind: T) {
  return ExtensionsStore[kind] as ExtensionsOfKind<T> | undefined;
}

export function registerExtension<T extends ExtensionKind>(extension: Extension<T>) {
  const record = (ExtensionsStore[extension.kind] || {}) as ExtensionsOfKind<T>;
  record[extension.name] = extension as Extension<T>;
  ExtensionsStore[extension.kind] = record as MappedRecord<ExtensionKind>[T];
}
