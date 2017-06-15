#include <memory.h>
#include <stdio.h>
#include <string.h>
#include <libjsonnet.h>
#include "_cgo_export.h"

char* CallImport_cgo(void *ctx, const char *base, const char *rel, char **found_here, int *success) {
  struct JsonnetVm* vm = ctx;
  return go_call_import(vm, (char*)base, (char*)rel, found_here, success);
}
}
