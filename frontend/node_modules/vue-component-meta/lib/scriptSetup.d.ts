import * as core from '@vue/language-core';
import type * as ts from 'typescript';
export declare function getDefaultsFromScriptSetup(ts: typeof import('typescript'), printer: ts.Printer, language: core.Language<string>, componentPath: string): Map<string, string> | undefined;
export declare function resolveDefaultOptionExpression(ts: typeof import('typescript'), _default: ts.Expression): ts.Expression;
