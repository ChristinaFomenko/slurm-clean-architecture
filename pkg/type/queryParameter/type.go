package queryParameter

import (
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/type/pagination"
	"github.com/ChristinaFomenko/slurm-clean-architecture/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
