package service

import (
	"testing"
	"time"
)

func TestCalculateStreak(t *testing.T) {
	now := time.Date(2026, time.August, 15, 14, 0, 0, 0, time.UTC)
	tests := []struct {
		name           string
		executionsDate []time.Time
		now            time.Time
		want           int64
	}{
		{
			name:           "sem execuções deve retornar streak 0",
			executionsDate: []time.Time{},
			now:            now,
			want:           0,
		},
		{
			name: "executado apenas hoje deve retornar streak 1",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 15, 9, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 1,
		},
		{
			name: "executado apenas ontem deve manter o streak em 1 (ainda pode fazer hoje)",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 14, 18, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 1,
		},
		{
			name: "executado anteontem, ontem e hoje deve retornar streak 3",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 13, 12, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 14, 20, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 3,
		},
		{
			name: "streak quebrado (última execução há 2 dias)",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 13, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 0,
		},
		{
			name: "streak recente com buraco no passado (deve contar apenas a sequência atual)",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 11, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				// dia 13 faltou
				time.Date(2026, time.August, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 2,
		},
		{
			name: "streak recente com buraco no passado (deve contar apenas a sequência atual 2)",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 12, 10, 0, 0, 0, time.UTC),
				// dia 13 faltou
				time.Date(2026, time.August, 14, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 2,
		},
		{
			name: "múltiplas execuções no mesmo dia não devem duplicar a contagem do streak",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 14, 9, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 10, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 18, 0, 0, 0, time.UTC), // 2ª execução no mesmo dia
			},
			now:  now,
			want: 2,
		},
		{
			name: "streak longo estendido",
			executionsDate: []time.Time{
				time.Date(2026, time.August, 10, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 11, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 12, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 13, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 14, 8, 0, 0, 0, time.UTC),
				time.Date(2026, time.August, 15, 8, 0, 0, 0, time.UTC),
			},
			now:  now,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateStreak(tt.executionsDate, tt.now)
			if got != tt.want {
				t.Errorf("CalculateStreak() = %d; esperava %d", got, tt.want)
			}
		})
	}
}
