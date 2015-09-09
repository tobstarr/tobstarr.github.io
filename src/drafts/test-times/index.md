# Test Times

## Examnple

{ for i in 1 2 3 4; do time bundle exec rspec spec/services/search/elastic_search_adapter_spec.rb:102; done } 2>&1 | tee build.log
