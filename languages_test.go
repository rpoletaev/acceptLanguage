package acceptLanguage

import (
  "testing"
  "context"
)

func TestContext(t *testing.T) {
  languages := ParseLanguagesList("ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
  ctx := WithContext(context.Background(), languages)
  t.Log(Get(ctx))
}