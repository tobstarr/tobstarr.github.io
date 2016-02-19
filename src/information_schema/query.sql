
select data_type, CHARACTER_MAXIMUM_LENGTH, count(1)
from information_schema.columns
where table_schema = 'phraseapp_production'
group by data_type, CHARACTER_MAXIMUM_LENGTH
order by count(1);
