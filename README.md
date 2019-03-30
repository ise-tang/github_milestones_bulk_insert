# GitHub MileStones Bulk Insert

This is a command line tool for create GitHub Milestones from a CSV file.

## Usage

At first, You have to generate your personel token.  
Read [GitHub Help](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line).  
If you want to create milestones for private repository, Select repo grant.

Next, create `settings.yml`

```
owner: "<OWENER NAME>"
repo: "<REPOSITORY NAME>"
access_token: "<YOUR PERSONAL ACCESS TOKEN HERE>
```

And prepare csv file for milestones 
It have to include the header.

Date format is `yyyy/mm/dd`

```
title,description,due_on
sprint1,foo,2019/04/02
sprint2,bar,2019/04/09
```

Then execute.

```
go run src/milestone_bulk_insert.go milestones.csv
```




## License

BSD 3-Clause Licens