package dao

import (
    "log"
    "{{.PackageBasePath}}/do"
    "{{.PackageBasePath}}/entity"

    "github.com/jmoiron/sqlx"
)
//{{.Remark}}
type {{.ObjectName}}Dao struct {
    db *sqlx.DB
}

func New{{.ObjectName}}Dao(db *sqlx.DB) *{{.ObjectName}}Dao {
    return &{{.ObjectName}}Dao {
        db: db,
    }
}

// 插入
func (o *{{.ObjectName}}Dao) Insert(p *do.{{.ObjectName}}) error {

    sql := `
       insert into {{.TableName}} ({{GenColumnNames .Fields}}) 
       values({{GenFieldHolders .Fields}})
    `
    log.Printf("EXEC SQL:%s , %v", sql, p)
    _, err := o.db.NamedExec(sql, p)

    return err

}

// 按条件更新
func (o *{{.ObjectName}}Dao) UpdateByPK(p *do.{{.ObjectName}}) error {
    sql := `
        update {{.TableName}} 
        set 
          {{ range $i,$v := .Fields }} {{GenUpdateStmt $i $v}} {{- end }}
        where 
          {{ with .PrimaryKey }}
          {{- range $i,$v := .}} {{GenPKWhere $i $v}} {{- end }}
          {{- else}}  1=2 {{- end}}
    `
    log.Printf("EXEC SQL:%s , %v", sql, p)
    _, err := o.db.NamedExec(sql, p)

    return err

}

// 按主键删除
func (o *{{.ObjectName}}Dao) DeleteByPK(p *do.{{.ObjectName}}) error {
    sql := `
       delete from {{.TableName}} 
       where 
          {{ with .PrimaryKey }}
          {{- range $i,$v := .}} {{GenPKWhere $i $v}} {{- end }}
          {{- else}}  1=2 {{- end}}
    `
    log.Printf("EXEC SQL:%s , %v", sql, p)
    _, err := o.db.NamedExec(sql, p)

    return err
}

// 按条件查询数据列表
func (o *{{.ObjectName}}Dao) Select(p *do.{{.ObjectName}}) ([]entity.{{.ObjectName}}, error) {
    sql := `
        select
            {{GenColumnNames .Fields}}
        from {{.TableName}}
        where
            1=1
    `
    {{- range $i,$v := .Fields}}
    if p.{{$v.FieldName}} != nil {
        sql += " and {{$v.ColumnName}}=:{{$v.ColumnName}} "
    }
    {{- end}}

    log.Printf("EXEC SQL:%s , %v", sql, p)
    rows, err := o.db.NamedQuery(sql, p)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    list := []entity.{{.ObjectName}}{}
    u := entity.{{.ObjectName}}{}
    for rows.Next() {
        err := rows.StructScan(&u)
        if err != nil {
            return nil, err
        }
        list = append(list, u)
    }

    return list, nil
}

// 按主键查询
func (o *{{.ObjectName}}Dao) GetByPK(p *do.{{.ObjectName}}) (*entity.{{.ObjectName}}, error) {
    sql := `
    select 
          {{GenColumnNames .Fields}}
    from  {{.TableName}}
    where 
          {{ with .PrimaryKey }}
          {{- range $i,$v := .}} {{GenPKWhere $i $v}} {{- end }}
          {{- else}}  1=2 {{- end}}
    `
    log.Printf("EXEC SQL:%s , %v", sql, p)
    rows, err := o.db.NamedQuery(sql, p)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    u := entity.{{.ObjectName}}{}
    if rows.Next() {
        err := rows.StructScan(&u)
        if err != nil {
            return nil, err
        }
    }

    return &u, nil
}
